#version 130

out float depth;

void main() {
    depth = gl_FragCoord.z;
}
