package shaders

var CRT = []byte(`
package main

var Time float

func curveRemapUV(uv vec2) vec2 {
	ret_uv := uv * 2.0-1.0
	offset := abs(vec2(ret_uv.y, ret_uv.x)) / vec2(5.0, 5.0)
	ret_uv = ret_uv + ret_uv * offset * offset
	ret_uv = ret_uv * 0.5 + 0.5
	return ret_uv
}

func scanLineIntensity(uv float, resolution float, opacity float) vec4 {
	pi := 3.1415926538
	intensity := sin(uv * resolution * pi * 2.0)
	intensity = ((0.5 * intensity) + 0.5) * 0.9 + 0.1
	return vec4(vec3(pow(intensity, opacity)), 1.0)
}

func vignetteIntensity(uv vec2, resolution vec2, opacity float, roundness float) vec4 {
	intensity := uv.x * uv.y * (1.0 - uv.x) * (1.0 - uv.y)
	return vec4(vec3(clamp(pow((resolution.x / roundness) * intensity, opacity), 0.0, 1.0)), 1.0);
}

func get_color_scanline(uv vec2, resolution vec2, c vec4, time float, lines_velocity float, lines_distance float, scanline_alpha float, scan_size float) vec4 {
	line_row := floor((uv.y * resolution.y/scan_size) + mod(Time*lines_velocity, lines_distance))
	n := 1.0 - ceil((mod(line_row,lines_distance)/lines_distance))
	ret := c - n*c*(1.0 - scanline_alpha)
	ret.a = 1.0
	return ret
}

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	remappedUV := curveRemapUV(texCoord)
	base_color := imageSrc0UnsafeAt(remappedUV)
	img_dst_tex_size := imageDstTextureSize()

	if (remappedUV.x < 0.0 || remappedUV.y < 0.0 || remappedUV.x > 1.0 || remappedUV.y > 1.0) {
		base_color = vec4(0.0, 0.0, 0.0, 1.0)
	}

	base_color = base_color * vignetteIntensity(remappedUV, img_dst_tex_size, 2.0, 0.1)

	base_color = base_color * scanLineIntensity(remappedUV.x, img_dst_tex_size.y, 0.05)
	base_color = base_color * scanLineIntensity(remappedUV.y, img_dst_tex_size.x, 0.05)

	// increase brightness
	base_color = base_color * vec4(vec3(1.1), 1.0)

	return get_color_scanline(remappedUV, img_dst_tex_size, base_color, Time, 0.03, 4.0, 0.9, 10.0)
}
`)
