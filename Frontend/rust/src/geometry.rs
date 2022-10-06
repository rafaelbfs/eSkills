use std::convert::TryInto;
use std::f32::consts::PI;
use std::ops::{self, Mul, MulAssign};
use std::ops::{Deref, Index, IndexMut};

fn greatest<T : PartialOrd<T> + Copy>(nums: &[T]) -> T {
    let mut max = nums[0];
    for i in 1..nums.len() {
        max = if max > nums[i] {max} else {nums[i]}
    }
    max
}

fn mk_face(verts: [Vertex; 4]) -> [f32; 18]
{
    let mut res: [f32; 18] = [0.0; 18];
    for i in (0..res.len()).step_by(3) {
        let v = if i < 9 {i / 3} else {6 - (i / 3)};
        res[i] = verts[v].x;
        res[i + 1] = verts[v].y;
        res[i + 2] = verts[v].z;
    }
    res
}

fn assign_face(dst: &mut [f32; 108], start: usize, verts: [f32; 18])
{
    for i in 0..verts.len() {
       dst[start + i] = verts[i];
    }
}

enum Axle {
    X, Y, Z
}

#[derive(Copy, Clone)]
#[derive(Debug)]
struct Vertex {
    x: f32,
    y: f32,
    z: f32
}

impl std::fmt::Display for Vertex {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "Vertex({}, {}, {})", self.x, self.y, self.z)
    }
}

impl Vertex {
    fn new(x: f32, y: f32, z: f32) -> Self {
        Vertex{x, y, z}
    }
    fn from_slice(v: &[f32]) -> Self {
        Vertex{x: v[0], y: v[1], z: v[2]}
    }
    fn zero() -> Self {
        Vertex::new(0.0, 0.0, 0.0)
    }
    fn cross(&self, other: &Vertex) -> Vertex {
        Vertex{x: self.y * other.z - self.z * other.y,
            y: self.z * other.x - self.x * other.z,
            z: self.x * other.y - self.z * other.x
        }

    }
}

impl ops::Sub<Vertex> for Vertex {
    type Output = Vertex;

    fn sub(self, rhs: Vertex) -> Self::Output {
        Vertex{x: self.x - rhs.x, y: self.y - rhs.y, z: self.z - rhs.z}
    }
}

impl ops::Mul<Vertex> for Vertex {
    type Output = Self;

    fn mul(self, rhs: Vertex) -> Self::Output {
        self.cross(&rhs)
    }
}

pub fn generate_cube() -> [f32; 108] {
    generate_box(1.0, 1.0, 1.0)
}

pub fn generate_box(length: f32, height: f32, depth: f32) -> [f32; 108] {
    let normal = greatest(&[1.0, length, height, depth]);
    let origin: f32 = -0.5;
    let mut vertices: [Vertex; 8] = [Vertex {x: origin, y: origin, z: origin}; 8];
    let triangles = &mut [0.0 as f32; 108];
    for i in 0..vertices.len() {
        let plus_x: f32 = ((i % 2) as f32) * length / normal;
        let plus_y: f32 = (((i >> 1) & 1) as f32) * height / normal;
        let plus_z: f32 = (((i >> 2) & 1) as f32) * depth / normal;
        vertices[i].x += plus_x;
        vertices[i].y += plus_y;
        vertices[i].z += plus_z;
    }

    for i in 0..6 {
        let bit = 2 - i / 2;
        let face1: [Vertex; 4] = (0..8).filter(|n| (n >> bit) % 2 == i % 2)
            .map(|n| vertices[n])    .collect::<Vec<Vertex>>()
            .try_into()
            .unwrap();
        assign_face(triangles, (i * 18), mk_face(face1));
    }

    *triangles
}

fn assign_triangle(buffer: &mut [f32], offset: usize, vertices: &[Vertex; 3]) {
    for i in 0..vertices.len() {
        buffer[offset + i * 3] = vertices[i].x;
        buffer[offset + i * 3 + 1] = vertices[i].y;
        buffer[offset + i * 3 + 2] = vertices[i].z;
    }
}

pub fn generate_tetrahedron(height: f32) -> [f32; 36] {
    let result = &mut [0.0 as f32; 36];
    let centroid = Vertex::zero();
    let top_y = height/2.0;
    let top = Vertex::new(0.0, top_y, 0.0);
    let cos_120deg = 0.5;
    let sin_60deg = js_sys::Math::sin((PI/3.0).into()) as f32;
    let a = Vertex::new(top_y, -top_y, 0.0);
    let b = Vertex::new(-cos_120deg * top_y, -top_y, sin_60deg * top_y);
    let c = Vertex::new(-cos_120deg * top_y, -top_y, -sin_60deg * top_y);

    assign_triangle(result, 0, &[a, b, c]);
    assign_triangle(result, 9, &[a, b, top]);
    assign_triangle(result, 9 * 2, &[b, c, top]);
    assign_triangle(result, 9 * 3, &[c, a, top]);

    *result
}

#[derive(Debug)]
pub struct Matrix4x4 {
    mat: [f32; 16]
}

impl std::fmt::Display for Matrix4x4 {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        let mut b = String::new();
        for i in (0..4) {
            b.push_str(format!("[{} {} {} {}]\n", self[(i, 0)], self[(i, 1)], self[(i, 2)], self[(i, 3)]).as_str());
            
        }
        
        write!(f, "{}", b)
    }
}

impl   Matrix4x4 {
    
    pub fn new (n: f32) -> Self {
        let mat: [f32; 16] = [n; 16];
        Matrix4x4{mat}
    }
    pub fn rotation_on_x(alpha: f32) -> Self {
        let cos_alpha = js_sys::Math::cos(alpha.into()) as f32;
        let sin_alpha = js_sys::Math::sin(alpha.into()) as f32;
        let mut m: Matrix4x4 = Matrix4x4::new(0.0);
        m[(0, 0)] = 1.0.into();
        m[(3, 3)] = 1.0.into();
        m[(1, 1)] = cos_alpha;
        m[(1, 2)] = -sin_alpha;
        m[(2, 2)] = cos_alpha;
        m[(2, 1)] = sin_alpha;

        m
    }
    pub fn rotation_on_y(alpha: f32) -> Self {
        let cos_alpha = js_sys::Math::cos(alpha.into()) as f32;
        let sin_alpha = js_sys::Math::sin(alpha.into()) as f32;
        let mut m: Matrix4x4 = Matrix4x4::new(0.0);
        m[(1, 1)] = 1.0;
        m[(3, 3)] = 1.0;
        m[(0, 0)] = cos_alpha;
        m[(0, 2)] = -sin_alpha;
        m[(2, 2)] = cos_alpha;
        m[(2, 0)] = sin_alpha;

        m
    }
    pub fn rotation_on_z(alpha: f32) -> Self {
        let cos_alpha = js_sys::Math::cos(alpha.into()) as f32;
        let sin_alpha = js_sys::Math::sin(alpha.into()) as f32;
        let mut m: Matrix4x4 = Matrix4x4::new(0.0);
        m[(2, 2)] = 1.0;
        m[(3, 3)] = 1.0;
        m[(0, 0)] = cos_alpha;
        m[(0, 1)] = -sin_alpha;
        m[(1, 1)] = cos_alpha;
        m[(1, 0)] = sin_alpha;

        m
    }
}
impl Into<[f32; 16]> for Matrix4x4 {
    fn into(self) -> [f32; 16] {        
        self.mat
     }
}
impl Clone for Matrix4x4 {
    fn clone(&self) -> Self {
        Matrix4x4{mat: self.mat.clone()}
    }
} 
impl Index<(usize, usize)> for Matrix4x4 {
    type Output = f32;

    fn index(&self, index: (usize, usize)) -> &Self::Output {
        &self.mat[index.0 * 4 + index.1]
    }
}

impl IndexMut<(usize, usize)> for Matrix4x4 {
    fn index_mut(&mut self, index: (usize, usize)) -> &mut Self::Output {
        &mut self.mat[index.0 * 4 + index.1]
    }
}

impl Mul for Matrix4x4 
{
    type Output = Self;
    
    fn mul(self, rhs: Self) -> Self
    {
        let mut res = Matrix4x4::new(0.0);
        for i in 0..4 {
            for j in 0..4 {
                let mut acc = 0.0;
                for k in 0..4 {
                    acc += self[(i,k)] * rhs[(k,j)];
                }
                res[(i,j)] = acc;
            }
        }
        res
    }
}

impl MulAssign for Matrix4x4 
{
    fn mul_assign(&mut self, rhs: Self)
    {
        let old = self.clone();
        for i in 0..4 {
            for j in 0..4 {
                for k in 0..4 {
                    self[(i,j)] += old[(i,k)] * rhs[(k,j)];
                }
            }
        }
    }
}

pub fn mk_xrotation_mat(alpha: f32) -> [f32; 16] {
    Matrix4x4::rotation_on_x(alpha).mat
}

pub fn mk_yrotation_mat(theta: f32) ->  [f32; 16] {
    Matrix4x4::rotation_on_y(theta).mat
}

pub fn mk_zrotation_mat(phi: f32) ->  [f32; 16] {
    Matrix4x4::rotation_on_z(phi).mat
}

pub fn calculate_normals(vertices: &[f32]) -> Vec<f32> {
    assert!(vertices.len() % 9 == 0, "buffer does not contain triangles");
    let mut normals = Vec::with_capacity(vertices.len());
    for i in (0..vertices.len()).step_by(9) {
        let a = Vertex::from_slice(&vertices[i..i+3]);
        let b = Vertex::from_slice(&vertices[i+3..i+6]);
        let c = Vertex::from_slice(&vertices[i+6..i+9]);
        let normal = (b - a) * (c - a);
        for _j in (0..9).step_by(3) {
            normals.push(normal.x);
            normals.push(normal.y);
            normals.push(normal.z);
        }
    }
    normals
}


