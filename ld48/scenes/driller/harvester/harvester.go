components {
  id: "ai"
  component: "/ld48/scenes/driller/harvester/harvester_ai.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "shadow_caster"
  component: "/ld48/scenes/common/shadow_caster.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "color"
    value: "0.0, 1.0, 1.0, 1.0"
    type: PROPERTY_TYPE_VECTOR4
  }
  properties {
    id: "max_distance"
    value: "-500.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "radius"
    value: "250.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "trunk1"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/brown_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"trunk1\"\n"
  ""
  position {
    x: 0.0
    y: 50.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "trunk2"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/brown_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"trunk2\"\n"
  ""
  position {
    x: 0.0
    y: 150.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "leaves_center"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/green_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"leaves_center\"\n"
  ""
  position {
    x: 0.0
    y: 250.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "leaves_n"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/green_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"leaves_n\"\n"
  ""
  position {
    x: 0.0
    y: 250.0
    z: 100.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "leaves_s"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/green_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"leaves_s\"\n"
  ""
  position {
    x: 0.0
    y: 250.0
    z: -100.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "leaves_e"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/green_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"leaves_e\"\n"
  ""
  position {
    x: 100.0
    y: 250.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "leaves_w"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/green_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"leaves_w\"\n"
  ""
  position {
    x: -100.0
    y: 250.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "leaves_top"
  type: "model"
  data: "mesh: \"/ld48/assets/greyboxes/cube_100.dae\"\n"
  "material: \"/ld48/materials/dirlight/dirlight_model_global_ambient.material\"\n"
  "textures: \"/ld48/assets/greyboxes/green_128.png\"\n"
  "skeleton: \"\"\n"
  "animations: \"\"\n"
  "default_animation: \"\"\n"
  "name: \"leaves_top\"\n"
  ""
  position {
    x: 0.0
    y: 350.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "phys_colobj"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 1.0\n"
  "restitution: 0.1\n"
  "group: \"body\"\n"
  "mask: \"plant\"\n"
  "mask: \"body\"\n"
  "mask: \"floor\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 150.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 50.0\n"
  "  data: 200.0\n"
  "  data: 50.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "shoot_colobj"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"harvester\"\n"
  "mask: \"harvester\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 200.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 150.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
