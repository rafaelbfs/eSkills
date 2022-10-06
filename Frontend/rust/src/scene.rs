use web_sys::{ WebGl2RenderingContext, WebGlBuffer, WebGlProgram, WebGlShader, WebGlUniformLocation, WebGlVertexArrayObject};

pub struct WebGLWrapper {
    ctx: WebGl2RenderingContext,
    program: WebGlProgram,
    vao: WebGlVertexArrayObject,
}

impl WebGLWrapper {
    pub fn new(ctx: WebGl2RenderingContext, program: WebGlProgram) -> Result<Self, String> {
        let vao = ctx
            .create_vertex_array()
            .ok_or_else(|| String::from("Could not create vertex array object"))?;
        ctx.bind_vertex_array(Some(&vao));
        Ok(Self { ctx, program, vao })
    }
    
    pub fn use_program(&self) {
        self.ctx.use_program(Some(&self.program));
    }
    
    pub fn bind_vao(&self) {
        self.ctx.bind_vertex_array(Some(&self.vao));
    }

    pub fn create_buffer(&self, attr: &str) -> Result<(i32, WebGlBuffer), String> {
        let location = self.ctx.get_attrib_location(&self.program, attr);
        let buffer = self.ctx.create_buffer().ok_or("Failed to create buffer")?;
        self.ctx
            .bind_buffer(WebGl2RenderingContext::ARRAY_BUFFER, Some(&buffer));
        Ok((location, buffer))
    }

    pub fn bind_and_fill_buffer(&self, location: i32, buffer: &WebGlBuffer, data: &[f32]) {
        self.ctx
            .bind_buffer(WebGl2RenderingContext::ARRAY_BUFFER, Some(&buffer));
        // Note that `Float32Array::view` is somewhat dangerous (hence the
        // `unsafe`!). This is creating a raw view into our module's
        // `WebAssembly.Memory` buffer, but if we allocate more pages for ourself
        // (aka do a memory allocation in Rust) it'll cause the buffer to change,
        // causing the `Float32Array` to be invalid.
        //
        // As a result, after `Float32Array::view` we have to be very careful not to
        // do any memory allocations before it's dropped.
        unsafe {
            let positions_array_buf_view = js_sys::Float32Array::view(data);

            self.ctx.buffer_data_with_array_buffer_view(
                WebGl2RenderingContext::ARRAY_BUFFER,
                &positions_array_buf_view,
                WebGl2RenderingContext::STATIC_DRAW,
            );
        }

        self.ctx.enable_vertex_attrib_array(location as u32);

        self.ctx.vertex_attrib_pointer_with_i32(
            location as u32,
            3,
            WebGl2RenderingContext::FLOAT,
            false,
            0,
            0,
        );

        self.ctx.bind_vertex_array(Some(&self.vao));
    }

    pub fn set_uniform_with_m4_of_f32(&self, loc: &WebGlUniformLocation, val: [f32; 16]) {
        self.ctx
            .uniform_matrix4fv_with_f32_array(Some(loc), false, &val);
    }

    pub fn set_uniform_by_name_with_m4_of_f32(
        &self,
        name: &str,
        val: [f32; 16],
    ) -> Result<WebGlUniformLocation, String> {
        let loc = self .ctx .get_uniform_location(&self.program, name)
            .ok_or_else(|| String::from(format!("No uniform named: {}", name)))?;
        self.set_uniform_with_m4_of_f32(&loc, val);
        Ok(loc)
    }
    
    pub fn clear(&self) {
        self.ctx.clear_color(0.0, 0.0, 0.0, 1.0);
        self.ctx.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);
    }

    pub fn draw(&self, vert_count: i32) {
        self.ctx
            .draw_arrays(WebGl2RenderingContext::TRIANGLES, 0, vert_count);
    }
}

pub fn compile_shader(
    context: &WebGl2RenderingContext,
    shader_type: u32,
    source: &str,
) -> Result<WebGlShader, String> {
    let shader = context
        .create_shader(shader_type)
        .ok_or_else(|| String::from("Unable to create shader object"))?;
    context.shader_source(&shader, source);
    context.compile_shader(&shader);

    if context
        .get_shader_parameter(&shader, WebGl2RenderingContext::COMPILE_STATUS)
        .as_bool()
        .unwrap_or(false)
    {
        Ok(shader)
    } else {
        Err(context
            .get_shader_info_log(&shader)
            .unwrap_or_else(|| String::from("Unknown error creating shader")))
    }
}

pub fn link_program(
    context: &WebGl2RenderingContext,
    vert_shader: &WebGlShader,
    frag_shader: &WebGlShader,
) -> Result<WebGlProgram, String> {
    let program = context
        .create_program()
        .ok_or_else(|| String::from("Unable to link program"))?;

    context.attach_shader(&program, vert_shader);
    context.attach_shader(&program, frag_shader);
    context.link_program(&program);

    if context
        .get_program_parameter(&program, WebGl2RenderingContext::LINK_STATUS)
        .as_bool()
        .unwrap_or(false)
    {
        Ok(program)
    } else {
        Err(context
            .get_program_info_log(&program)
            .unwrap_or_else(|| String::from("Unknown error creating program object")))
    }
}
