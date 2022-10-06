extern crate core;

mod geometry;
mod scene;
mod utils;

use crate::geometry::{
    calculate_normals, generate_cube, generate_tetrahedron, mk_xrotation_mat, mk_yrotation_mat,
    mk_zrotation_mat, Matrix4x4
};
use crate::scene::{compile_shader, link_program, WebGLWrapper};
use crate::utils::set_panic_hook;
use wasm_bindgen::prelude::*;
use wasm_bindgen::JsCast;
use web_sys::{WebGl2RenderingContext, WebGlProgram, WebGlShader, WebGlBuffer};
use std::cell::RefCell;
use std::rc::Rc;
use web_sys::console;

const FULL_CIRCLE: f32 = 2.0 * std::f32::consts::PI;

struct State {
    vertices: Vec<f32>,
    vert_buff: (i32, WebGlBuffer),
    normals: Vec<f32>,
    normal_buff: (i32, WebGlBuffer),
    rot_x: Matrix4x4,
    rot_y: Matrix4x4,
    rot_z: Matrix4x4,
    omega_x: f32,
    omega_y: f32,
    omega_z: f32,
    alpha: f32,
    theta: f32,
    phi: f32,
    timestamp: f64 
}

fn normalize_angle(angle: f32) -> f32 {
     if angle > FULL_CIRCLE {
        return angle - FULL_CIRCLE;
    }
    angle
}

/* impl State {
    fn update(&mut self, now_ts: f64, do_print: bool) -> Result<(), String> {
        self.ctx.use_program();        
        let delta: f32 = (now_ts - self.timestamp) as f32;
        
        if do_print {
//            console::log_1(&format!("X Rotation\n{}", self.rot_x).into());
            console::log_1(&format!("Y Rotation\n{}", self.rot_y).into());
        }
        
        self.timestamp = now_ts;
        self.ctx.bind_and_fill_buffer(self.vert_buff.0, &self.vert_buff.1, self.vertices.as_slice());
        
        if (self.omega_x != 0.0) {
            self.alpha = normalize_angle(self.alpha + self.omega_x * delta);
            self.rot_x = Matrix4x4::rotation_on_x(self.alpha);
            self.ctx.set_uniform_by_name_with_m4_of_f32("rotationX", self.rot_x.clone().into())?;
        }
        if (self.omega_y != 0.0) {
            self.theta = normalize_angle(self.theta + self.omega_y * delta);
            self.rot_y = Matrix4x4::rotation_on_y(self.theta);
            self.ctx.set_uniform_by_name_with_m4_of_f32("rotationY", self.rot_y.clone().into())?;
        }
        if (self.omega_z != 0.0) {
            self.phi = normalize_angle(self.phi + self.omega_z * delta);
            self.rot_z = Matrix4x4::rotation_on_z(self.phi);
            self.ctx.set_uniform_by_name_with_m4_of_f32("rotationZ", self.rot_z.clone().into())?;
        }
        
        self.ctx.bind_vao();
        
        Ok(())        
    }
} */

fn window() -> web_sys::Window {
    web_sys::window().expect("no global `window` exists")
}

fn performance() -> web_sys::Performance {
    window().performance().expect("Performance obj is required for the correct function of this program")
} 

fn request_animation_frame(f: &Closure<dyn FnMut()>) {
    window()
        .request_animation_frame(f.as_ref().unchecked_ref())
        .expect("should register `requestAnimationFrame` OK");
}

#[wasm_bindgen(start)]
pub fn start() -> Result<(), JsValue> {
    set_panic_hook();
    let document = web_sys::window().unwrap().document().unwrap();
    let canvas = document.get_element_by_id("canvas").unwrap();
    let canvas: web_sys::HtmlCanvasElement = canvas.dyn_into::<web_sys::HtmlCanvasElement>()?;

    let context = canvas
        .get_context("webgl2")?
        .unwrap()
        .dyn_into::<WebGl2RenderingContext>()?;

    let vert_shader = compile_shader(
        &context,
        WebGl2RenderingContext::VERTEX_SHADER,
        r##"#version 300 es

        precision mediump float;
        in vec4 position;
        in vec3 normal;

        uniform mat4 rotationX;
        uniform mat4 rotationY;
        uniform mat4 rotationZ;

        out vec4 v_color;
        out vec3 v_normal;

        void main() {
            mat4 rotation = rotationX * rotationY * rotationZ;
            gl_Position = rotation * position;
            v_color = position * 0.5 + 0.5;
            v_normal = normal;
        }
        "##,
    )?;

    let frag_shader = compile_shader(
        &context,
        WebGl2RenderingContext::FRAGMENT_SHADER,
        r##"#version 300 es

        precision mediump float;
        in vec4 v_color;
        in vec3 v_normal;

        out vec4 outColor;

        uniform vec3 u_light;

        void main() {
            outColor = v_color;
        }
        "##,
    )?;
    let program = link_program(&context, &vert_shader, &frag_shader)?;
    context.use_program(Some(&program));

    let gl = WebGLWrapper::new(context, program)?;

    /* let mut vertices: [f32; 27] = [0.0; 27];

    for i in 0 .. vertices.len() {
        vertices[i] = (js_sys::Math::random() as f32) * 2.0 - 1.0;
      console::log_1(&format!("Element {} = {} ",i, vertices[i]).into());
    } */

    //let vertices = generate_cube();
    let vertices = generate_tetrahedron(1.0);
    let normals = calculate_normals(&vertices).to_owned();

    let (verts_loc, verts_buf) = gl.create_buffer("position")?;
    gl.bind_and_fill_buffer(verts_loc, &verts_buf, &vertices);
    
    let (norm_loc, norm_buf) = gl.create_buffer("normal")?;
    gl.bind_and_fill_buffer(norm_loc, &norm_buf, normals.as_slice());

    let vert_count = (vertices.len() / 3) as i32;
    

    let mut st = State {
        vertices: vertices.to_vec(),
        vert_buff: (verts_loc, verts_buf),
        normals,
        normal_buff: (norm_loc, norm_buf),
        omega_x: 0.0,
        omega_y: 0.25 * std::f32::consts::PI,
        omega_z: 0.0,
        timestamp: performance().now(),
        rot_x: Matrix4x4::rotation_on_x(0.005 * std::f32::consts::PI),
        rot_y: Matrix4x4::rotation_on_y(0.2 * std::f32::consts::PI),
        rot_z: Matrix4x4::rotation_on_z(0.0001 * std::f32::consts::PI),
        alpha: 0.0,
        theta: 0.001 * FULL_CIRCLE,
        phi: 0.1 * FULL_CIRCLE
    };
    /*
    st.update(performance().now(), true)?;
    //context.enable(WebGl2RenderingContext::CULL_FACE);
    let mut counter: u64 = 0;
    let f = Rc::new(RefCell::new(None));
    let g = f.clone();

    
    *g.borrow_mut() = Some(Closure::new(move || {
        let now = performance().now();
        //document.get_element_by_id("overhead").map(|el| {el.set_text_content(Some(format!("Now {}", now).as_str())); "OK"});
        st.ctx.clear();
        if let Err(s) = st.update(now, counter % 180 == 0) {
            
          console::log_1(&format!("Error while changing state: {} ", s).into());
          return;
        } 
        st.ctx.draw(vert_count);
        counter += 1;
        // Schedule ourself for another requestAnimationFrame callback.
        request_animation_frame(f.borrow().as_ref().unwrap());
    }));

    request_animation_frame(g.borrow().as_ref().unwrap()); */
    gl.set_uniform_by_name_with_m4_of_f32("rotationX", st.rot_x.into());
    gl.set_uniform_by_name_with_m4_of_f32("rotationY", st.rot_y.into());
    gl.set_uniform_by_name_with_m4_of_f32("rotationZ", st.rot_z.into());
    
    
    gl.clear();
    
    gl.draw(vert_count);
    
    Ok(())
}

