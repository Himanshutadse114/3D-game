components {
  id: "anim"
  component: "/ld48/scenes/common/rock_anim.script"
}
embedded_components {
  id: "model_main"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "name: \"main\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/ld48/assets/greyboxes/illustrative_rock.png\"\n"
  "  }\n"
  "}\n"
  ""
}
embedded_components {
  id: "model_jag1"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "name: \"jag1\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/ld48/assets/greyboxes/illustrative_rock.png\"\n"
  "  }\n"
  "}\n"
  ""
  position {
    x: 40.0
    y: -20.0
    z: 30.0
  }
  rotation {
    x: 0.2
    y: 0.5
    z: 0.1
  }
}
embedded_components {
  id: "model_jag2"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "name: \"jag2\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/ld48/assets/greyboxes/illustrative_rock.png\"\n"
  "  }\n"
  "}\n"
  ""
  position {
    x: -30.0
    y: -10.0
    z: -40.0
  }
  rotation {
    x: -0.3
    y: -0.4
    z: 0.2
  }
}
