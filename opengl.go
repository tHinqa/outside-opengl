// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package opengl

const (
	Flat   = 0x1D00
	Smooth = 0x1D01
)

type (
	fix uintptr

	GLcharARB   fix
	EglImageOES fix
	HandleARB   fix
	Intptr      ptrdiff_t
	IntptrARB   ptrdiff_t
	Sizeiptr    ptrdiff_t
	SizeiptrARB ptrdiff_t
	ptrdiff_t   fix
	// _cl_context           fix
	// _cl_event             fix
	// GLDEBUGPROCAMD        fix
	// GLDEBUGPROCARB        fix
	// GLprogramcallbackMESA fix
	// GLvdpauSurfaceNV      fix

	enum       int
	void       struct{}
	GLbitfield uint
	GLboolean  bool
	GLbyte     int8
	GLchar     int8
	Clampd     float64
	Clampf     float32
	HalfNV     uint16
	Int64EXT   int64
	Sizei      int
	Sync       *struct{}
	uint64EXT  uint64
	GLvoid     struct{}

	GLUnurbs      struct{}
	GLUquadric    struct{}
	GLUtesselator struct{}
)

type AttribMask GLbitfield

const (
	CurrentBit AttribMask = 1 << iota
	PointBit
	LineBit
	PolygonBit
	PolygonStippleBit
	PixelModeBit
	LightingBit
	FogBit
	DepthBufferBit
	AccumBufferBit
	StencilBufferBit
	ViewportBit
	TransformBit
	EnableBit
	ColorBufferBit
	HintBit
	EvalBit
	ListBit
	TextureBit
	ScissorBit
	AllAttribBits AttribMask = 0x000fffff
)

type String enum

const (
	Vendor String = iota + 0x1F00
	Renderer
	Version
	Extensions
)

type Funcptr func()

var Accum func(op Enum, value float32)
var ActiveProgramEXT func(program uint)

// var ActiveShaderProgram func(pipeline, program uint)
// var ActiveStencilFaceEXT func(face Enum)
var ActiveTexture func(texture Enum)
var ActiveTextureARB func(texture Enum)
var ActiveVaryingNV func(program uint, name string)
var AlphaFragmentOp1ATI func(op Enum, dst, dstMod, arg1, arg1Rep, arg1Mod uint)
var AlphaFragmentOp2ATI func(op Enum, dst, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod uint)
var AlphaFragmentOp3ATI func(op Enum, dst, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod, arg3, arg3Rep, arg3Mod uint)
var AlphaFunc func(func_ Enum, ref Clampf)

// var ApplyTextureEXT func(mode Enum)
var AreProgramsResidentNV func(n Sizei, programs *uint, residences *GLboolean) GLboolean
var AreTexturesResident func(n Sizei, textures *uint, residences *GLboolean) GLboolean
var AreTexturesResidentEXT func(n Sizei, textures *uint, residences *GLboolean) GLboolean
var ArrayElement func(i int)
var ArrayElementEXT func(i int)

// var ArrayObjectATI func(array Enum, size int, type_ Enum, stride Sizei, buffer, offset uint)
// var AsyncMarkerSGIX func(marker uint)
var AttachObjectARB func(containerObj, obj HandleARB)
var AttachShader func(program, shader uint)
var Begin func(mode Enum)
var BeginConditionalRender func(id uint, mode Enum)
var BeginConditionalRenderNV func(id uint, mode Enum)
var BeginFragmentShaderATI func()

// var BeginOcclusionQueryNV func(id uint)
// var BeginPerfMonitorAMD func(monitor uint)
var BeginQuery func(target Enum, id uint)
var BeginQueryARB func(target Enum, id uint)

// var BeginQueryIndexed func(target Enum, index, id uint)
var BeginTransformFeedback func(primitiveMode Enum)
var BeginTransformFeedbackEXT func(primitiveMode Enum)

// var BeginTransformFeedbackNV func(primitiveMode Enum)
// var BeginVertexShaderEXT func()
// var BeginVideoCaptureNV func(video_capture_slot uint)
var BindAttribLocation func(program, index uint, name string)
var BindAttribLocationARB func(programObj HandleARB, index uint, name *GLcharARB)
var BindBuffer func(target Enum, buffer uint)
var BindBufferARB func(target Enum, buffer uint)
var BindBufferBase func(target Enum, index, buffer uint)
var BindBufferBaseEXT func(target Enum, index, buffer uint)

// var BindBufferBaseNV func(target Enum, index, buffer uint)
var BindBufferOffsetEXT func(target Enum, index, buffer uint, offset Intptr)

// var BindBufferOffsetNV func(target Enum, index, buffer uint, offset Intptr)
var BindBufferRange func(target Enum, index, buffer uint, offset Intptr, size Sizeiptr)
var BindBufferRangeEXT func(target Enum, index, buffer uint, offset Intptr, size Sizeiptr)

// var BindBufferRangeNV func(target Enum, index, buffer uint, offset Intptr, size Sizeiptr)
var BindFragDataLocation func(program, color uint, name string)
var BindFragDataLocationEXT func(program, color uint, name string)

// var BindFragDataLocationIndexed func(program, colorNumber, index uint, name string)
var BindFragmentShaderATI func(id uint)
var BindFramebuffer func(target Enum, framebuffer uint)
var BindFramebufferEXT func(target Enum, framebuffer uint)

// var BindImageTexture func(unit, texture uint, level int, layered GLboolean, layer int, access, format Enum)
// var BindImageTextureEXT func(index, texture uint, level int, layered GLboolean, layer int, access Enum, format int)
// var BindLightParameterEXT func(light, value Enum) uint
// var BindMaterialParameterEXT func(face, value Enum) uint
// var BindMultiTextureEXT func(texunit, target Enum, texture uint)
// var BindParameterEXT func(value Enum) uint
var BindProgramARB func(target Enum, program uint)
var BindProgramNV func(target Enum, id uint)

// var BindProgramPipeline func(pipeline uint)
var BindRenderbuffer func(target Enum, renderbuffer uint)
var BindRenderbufferEXT func(target Enum, renderbuffer uint)
var BindSampler func(unit, sampler uint)

// var BindTexGenParameterEXT func(unit, coord, value Enum) uint
var BindTexture func(target Enum, texture uint)
var BindTextureEXT func(target Enum, texture uint)

// var BindTextureUnitParameterEXT func(unit, value Enum) uint
var BindTransformFeedback func(target Enum, id uint)

// var BindTransformFeedbackNV func(target Enum, id uint)
var BindVertexArray func(array uint)

// var BindVertexArrayAPPLE func(array uint)
// var BindVertexShaderEXT func(id uint)
// var BindVideoCaptureStreamBufferNV func(video_capture_slot, stream uint, frame_region Enum, offset IntptrARB)
// var BindVideoCaptureStreamTextureNV func(video_capture_slot, stream uint, frame_region, target Enum, texture uint)
// var Binormal3bEXT func(bx, by, bz GLbyte)
// var Binormal3bvEXT func(v *GLbyte)
// var Binormal3dEXT func(bx, by, bz float64)
// var Binormal3dvEXT func(v *float64)
// var Binormal3fEXT func(bx, by, bz float32)
// var Binormal3fvEXT func(v *float32)
// var Binormal3iEXT func(bx, by, bz int)
// var Binormal3ivEXT func(v *int)
// var Binormal3sEXT func(bx, by, bz int16)
// var Binormal3svEXT func(v *int16)
// var BinormalPointerEXT func(type_ Enum, stride Sizei, pointer *GLvoid)
var Bitmap func(width, height Sizei, hxorig, yorig, xmove, ymove float32, bitmap *uint8)
var BlendColor func(red, green, blue, alpha Clampf)
var BlendColorEXT func(red, green, blue, alpha Clampf)
var BlendEquation func(mode Enum)
var BlendEquationEXT func(mode Enum)

// var BlendEquationi func(buf uint, mode Enum)
var BlendEquationiARB func(buf uint, mode Enum)
var BlendEquationIndexedAMD func(buf uint, mode Enum)
var BlendEquationSeparate func(modeRGB, modeAlpha Enum)

// var BlendEquationSeparateATI func(modeRGB, modeA Enum)
// var BlendEquationSeparateEXT func(modeRGB, modeAlpha Enum)
// var BlendEquationSeparatei func(buf uint, modeRGB, modeAlpha Enum)
var BlendEquationSeparateiARB func(buf uint, modeRGB, modeAlpha Enum)
var BlendEquationSeparateIndexedAMD func(buf uint, modeRGB, modeAlpha Enum)

var BlendFunc func(sfactor, dfactor Enum)

// var BlendFunci func(buf uint, src, dst Enum)
var BlendFunciARB func(buf uint, src, dst Enum)
var BlendFuncIndexedAMD func(buf uint, src, dst Enum)
var BlendFuncSeparate func(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum)
var BlendFuncSeparateEXT func(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum)

// var BlendFuncSeparatei func(buf uint, srcRGB, dstRGB, srcAlpha, dstAlpha Enum)
var BlendFuncSeparateiARB func(buf uint, srcRGB, dstRGB, srcAlpha, dstAlpha Enum)
var BlendFuncSeparateIndexedAMD func(buf uint, srcRGB, dstRGB, srcAlpha, dstAlpha Enum)

// var BlendFuncSeparateINGR func(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum)
var BlitFramebuffer func(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask GLbitfield, filter Enum)

// var BlitFramebufferEXT func(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask GLbitfield, filter Enum)
// var BufferAddressRangeNV func(pname Enum, index, address uint64EXT, length Sizeiptr)
var BufferData func(target Enum, size Sizeiptr, data *GLvoid, usage Enum)
var BufferDataARB func(target Enum, size SizeiptrARB, data *GLvoid, usage Enum)

// var BufferParameteriAPPLE func(target, pname Enum, param int)
var BufferSubData func(target Enum, offset Intptr, size Sizeiptr, data *GLvoid)
var BufferSubDataARB func(target Enum, offset IntptrARB, size SizeiptrARB, data *GLvoid)
var CallList func(list uint)
var CallLists func(n Sizei, type_ Enum, lists *GLvoid)
var CheckFramebufferStatus func(target Enum) Enum
var CheckFramebufferStatusEXT func(target Enum) Enum

// var CheckNamedFramebufferStatusEXT func(framebuffer uint, target Enum) Enum
var ClampColor func(target, clamp Enum)
var ClampColorARB func(target, clamp Enum)
var Clear func(mask AttribMask)
var ClearAccum func(hred, green, blue, alpha float32)
var ClearBufferfi func(buffer Enum, drawbuffer int, depth float32, stencil int)
var ClearBufferfv func(buffer Enum, drawbuffer int, value *float32)
var ClearBufferiv func(buffer Enum, drawbuffer int, value *int)
var ClearBufferuiv func(buffer Enum, drawbuffer int, value *uint)
var ClearColor func(hred, green, blue, alpha Clampf)
var ClearColorIiEXT func(red, green, blue, alpha int)
var ClearColorIuiEXT func(red, green, blue, alpha uint)

// var ClearDebugLogMESA func(obj HandleARB, logType, shaderType Enum)
var ClearDepth func(depth Clampd)

// var ClearDepthdNV func(depth float64)
var ClearDepthf func(d Clampf)
var ClearIndex func(c float32)
var ClearStencil func(s int)
var ClientActiveTexture func(texture Enum)
var ClientActiveTextureARB func(texture Enum)

// var ClientActiveVertexStreamATI func(stream Enum)
// var ClientAttribDefaultEXT func(mask GLbitfield)
var ClientWaitSync func(sync Sync, flags GLbitfield, timeout uint64) Enum
var ClipPlane func(plane Enum, equation *float64)
var Color3b func(hred, green, blue GLbyte)
var Color3bv func(v *GLbyte)
var Color3d func(hred, green, blue float64)
var Color3dv func(v *float64)
var Color3f func(hred, green, blue float32)
var Color3fv func(v *float32)

// var Color3fVertex3fSUN func(r, g, b, x, y, z float32)
// var Color3fVertex3fvSUN func(c, v *float32)
// var Color3hNV func(red, green, blue GLhalfNV)
// var Color3hvNV func(v *GLhalfNV)
var Color3i func(hred, green, blue int)
var Color3iv func(v *int)
var Color3s func(hred, green, blue int16)
var Color3sv func(v *int16)
var Color3ub func(hred, green, blue uint8)
var Color3ubv func(v *uint8)
var Color3ui func(hred, green, blue uint)
var Color3uiv func(v *uint)
var Color3us func(hred, green, blue uint16)
var Color3usv func(v *uint16)
var Color4b func(hred, green, blue, alpha GLbyte)
var Color4bv func(v *GLbyte)
var Color4d func(hred, green, blue, alpha float64)
var Color4dv func(v *float64)
var Color4f func(hred, green, blue, alpha float32)

// var Color4fNormal3fVertex3fSUN func(r, g, b, a, nx, ny, nz, x, y, z float32)
// var Color4fNormal3fVertex3fvSUN func(c, n, v *float32)
var Color4fv func(v *float32)

// var Color4hNV func(red, green, blue, alpha GLhalfNV)
// var Color4hvNV func(v *GLhalfNV)
var Color4i func(hred, green, blue, alpha int)
var Color4iv func(v *int)
var Color4s func(hred, green, blue, alpha int16)
var Color4sv func(v *int16)
var Color4ub func(hred, green, blue, alpha uint8)
var Color4ubv func(v *uint8)

// var Color4ubVertex2fSUN func(r, g, b, a uint8, x, y float32)
// var Color4ubVertex2fvSUN func(c *uint8, v *float32)
// var Color4ubVertex3fSUN func(r, g, b, a uint8, x, y, z float32)
// var Color4ubVertex3fvSUN func(c *uint8, v *float32)
var Color4ui func(hred, green, blue, alpha uint)
var Color4uiv func(v *uint)
var Color4us func(hred, green, blue, alpha uint16)
var Color4usv func(v *uint16)

// var ColorFormatNV func(size int, type_ Enum, stride Sizei)
var ColorFragmentOp1ATI func(op Enum, dst, dstMask, dstMod, arg1, arg1Rep, arg1Mod uint)
var ColorFragmentOp2ATI func(op Enum, dst, dstMask, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod uint)
var ColorFragmentOp3ATI func(op Enum, dst, dstMask, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod, arg3, arg3Rep, arg3Mod uint)
var ColorMask func(hred, green, blue, alpha GLboolean)
var ColorMaski func(index uint, r, g, b, a GLboolean)
var ColorMaskIndexedEXT func(index uint, r, g, b, a GLboolean)
var ColorMaterial func(face, mode Enum)
var ColorP3ui func(type_ Enum, color uint)
var ColorP3uiv func(type_ Enum, color *uint)
var ColorP4ui func(type_ Enum, color uint)
var ColorP4uiv func(type_ Enum, color *uint)
var ColorPointer func(size int, type_ Enum, stride Sizei, ptr *GLvoid)
var ColorPointerEXT func(size int, type_ Enum, stride, count Sizei, pointer *GLvoid)

// var ColorPointerListIBM func(size int, type_ Enum, stride int, pointer **GLvoid, ptrstride int)
// var ColorPointervINTEL func(size int, type_ Enum, pointer **GLvoid)
var ColorSubTable func(target Enum, start, count Sizei, format, type_ Enum, data *GLvoid)

// var ColorSubTableEXT func(target Enum, start, count Sizei, format, type_ Enum, data *GLvoid)
var ColorTable func(target, internalformat Enum, width Sizei, format, type_ Enum, table *GLvoid)
var ColorTableEXT func(target, internalFormat Enum, width Sizei, format, type_ Enum, table *GLvoid)
var ColorTableParameterfv func(target, pname Enum, params *float32)

// var ColorTableParameterfvSGI func(target, pname Enum, params *float32)
var ColorTableParameteriv func(target, pname Enum, params *int)

// var ColorTableParameterivSGI func(target, pname Enum, params *int)
// var ColorTableSGI func(target, internalformat Enum, width Sizei, format, type_ Enum, table *GLvoid)
var CombinerInputNV func(stage, portion, variable, input, mapping, componentUsage Enum)
var CombinerOutputNV func(stage, portion, abOutput, cdOutput, sumOutput, scale, bias Enum, abDotProduct, cdDotProduct, muxSum GLboolean)
var CombinerParameterfNV func(pname Enum, param float32)
var CombinerParameterfvNV func(pname Enum, params *float32)
var CombinerParameteriNV func(pname Enum, param int)
var CombinerParameterivNV func(pname Enum, params *int)

// var CombinerStageParameterfvNV func(stage, pname Enum, params *float32)
var CompileShader func(shader uint)
var CompileShaderARB func(shaderObj HandleARB)

// var CompileShaderIncludeARB func(shader uint, count Sizei, path []string, length *int)
// var CompressedMultiTexImage1DEXT func(texunit, target Enum, level int, internalformat Enum, width Sizei, border int, imageSize Sizei, bits *GLvoid)
// var CompressedMultiTexImage2DEXT func(texunit, target Enum, level int, internalformat Enum, width, height Sizei, border int, imageSize Sizei, bits *GLvoid)
// var CompressedMultiTexImage3DEXT func(texunit, target Enum, level int, internalformat Enum, width, height, depth Sizei, border int, imageSize Sizei, bits *GLvoid)
// var CompressedMultiTexSubImage1DEXT func(texunit, target Enum, level, xoffset int, width Sizei, format Enum, imageSize Sizei, bits *GLvoid)
// var CompressedMultiTexSubImage2DEXT func(texunit, target Enum, level, xoffset, yoffset int, width, height Sizei, format Enum, imageSize Sizei, bits *GLvoid)
// var CompressedMultiTexSubImage3DEXT func(texunit, target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format Enum, imageSize Sizei, bits *GLvoid)
var CompressedTexImage1D func(target Enum, level int, internalformat Enum, width Sizei, border int, imageSize Sizei, data *GLvoid)
var CompressedTexImage1DARB func(target Enum, level int, internalformat Enum, width Sizei, border int, imageSize Sizei, data *GLvoid)
var CompressedTexImage2D func(target Enum, level int, internalformat Enum, width, height Sizei, border int, imageSize Sizei, data *GLvoid)
var CompressedTexImage2DARB func(target Enum, level int, internalformat Enum, width, height Sizei, border int, imageSize Sizei, data *GLvoid)
var CompressedTexImage3D func(target Enum, level int, internalformat Enum, width, height, depth Sizei, border int, imageSize Sizei, data *GLvoid)
var CompressedTexImage3DARB func(target Enum, level int, internalformat Enum, width, height, depth Sizei, border int, imageSize Sizei, data *GLvoid)
var CompressedTexSubImage1D func(target Enum, level, xoffset int, width Sizei, format Enum, imageSize Sizei, data *GLvoid)
var CompressedTexSubImage1DARB func(target Enum, level, xoffset int, width Sizei, format Enum, imageSize Sizei, data *GLvoid)
var CompressedTexSubImage2D func(target Enum, level, xoffset, yoffset int, width, height Sizei, format Enum, imageSize Sizei, data *GLvoid)
var CompressedTexSubImage2DARB func(target Enum, level, xoffset, yoffset int, width, height Sizei, format Enum, imageSize Sizei, data *GLvoid)
var CompressedTexSubImage3D func(target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format Enum, imageSize Sizei, data *GLvoid)
var CompressedTexSubImage3DARB func(target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format Enum, imageSize Sizei, data *GLvoid)

// var CompressedTextureImage1DEXT func(texture uint, target Enum, level int, internalformat Enum, width Sizei, border int, imageSize Sizei, bits *GLvoid)
// var CompressedTextureImage2DEXT func(texture uint, target Enum, level int, internalformat Enum, width, height Sizei, border int, imageSize Sizei, bits *GLvoid)
// var CompressedTextureImage3DEXT func(texture uint, target Enum, level int, internalformat Enum, width, height, depth Sizei, border int, imageSize Sizei, bits *GLvoid)
// var CompressedTextureSubImage1DEXT func(texture uint, target Enum, level, xoffset int, width Sizei, format Enum, imageSize Sizei, bits *GLvoid)
// var CompressedTextureSubImage2DEXT func(texture uint, target Enum, level, xoffset, yoffset int, width, height Sizei, format Enum, imageSize Sizei, bits *GLvoid)
// var CompressedTextureSubImage3DEXT func(texture uint, target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format Enum, imageSize Sizei, bits *GLvoid)
var ConvolutionFilter1D func(target, internalformat Enum, width Sizei, format, type_ Enum, image *GLvoid)

// var ConvolutionFilter1DEXT func(target, internalformat Enum, width Sizei, format, type_ Enum, image *GLvoid)
var ConvolutionFilter2D func(target, internalformat Enum, width, height Sizei, format, type_ Enum, image *GLvoid)

// var ConvolutionFilter2DEXT func(target, internalformat Enum, width, height Sizei, format, type_ Enum, image *GLvoid)
var ConvolutionParameterf func(target, pname Enum, params float32)

// var ConvolutionParameterfEXT func(target, pname Enum, params float32)
var ConvolutionParameterfv func(target, pname Enum, params *float32)

// var ConvolutionParameterfvEXT func(target, pname Enum, params *float32)
var ConvolutionParameteri func(target, pname Enum, params int)

// var ConvolutionParameteriEXT func(target, pname Enum, params int)
var ConvolutionParameteriv func(target, pname Enum, params *int)

// var ConvolutionParameterivEXT func(target, pname Enum, params *int)
var CopyBufferSubData func(readTarget, writeTarget Enum, readOffset, writeOffset Intptr, size Sizeiptr)
var CopyColorSubTable func(target Enum, start Sizei, x, y int, width Sizei)

// var CopyColorSubTableEXT func(target Enum, start Sizei, x, y int, width Sizei)
var CopyColorTable func(target, internalformat Enum, x, y int, width Sizei)

// var CopyColorTableSGI func(target, internalformat Enum, x, y int, width Sizei)
var CopyConvolutionFilter1D func(target, internalformat Enum, x, y int, width Sizei)

// var CopyConvolutionFilter1DEXT func(target, internalformat Enum, x, y int, width Sizei)
var CopyConvolutionFilter2D func(target, internalformat Enum, x, y int, width, height Sizei)

// var CopyConvolutionFilter2DEXT func(target, internalformat Enum, x, y int, width, height Sizei)
// var CopyImageSubDataNV func(srcName uint, srcTarget Enum, srcLevel, srcX, srcY, srcZ int, dstName uint, dstTarget Enum, dstLevel, dstX, dstY, dstZ int, width, height, depth Sizei)
// var CopyMultiTexImage1DEXT func(texunit, target Enum, level int, internalformat Enum, x, y int, width Sizei, border int)
// var CopyMultiTexImage2DEXT func(texunit, target Enum, level int, internalformat Enum, x, y int, width, height Sizei, border int)
// var CopyMultiTexSubImage1DEXT func(texunit, target Enum, level, xoffset, x, y int, width Sizei)
// var CopyMultiTexSubImage2DEXT func(texunit, target Enum, level, xoffset, yoffset, x, y int, width, height Sizei)
// var CopyMultiTexSubImage3DEXT func(texunit, target Enum, level, xoffset, yoffset, zoffset, x, y int, width, height Sizei)
var CopyPixels func(x, y int, width, height Sizei, type_ Enum)
var CopyTexImage1D func(target Enum, level int, internalformat Enum, x, y int, width Sizei, border int)
var CopyTexImage1DEXT func(target Enum, level int, internalformat Enum, x, y int, width Sizei, border int)
var CopyTexImage2D func(target Enum, level int, internalformat Enum, x, y int, width, height Sizei, border int)
var CopyTexImage2DEXT func(target Enum, level int, internalformat Enum, x, y int, width, height Sizei, border int)
var CopyTexSubImage1D func(target Enum, hlevel, xoffset, x, y int, width Sizei)
var CopyTexSubImage1DEXT func(target Enum, level, xoffset, x, y int, width Sizei)
var CopyTexSubImage2D func(target Enum, hhlevel, xoffset, yoffset, x, y int, width, height Sizei)
var CopyTexSubImage2DEXT func(target Enum, level, xoffset, yoffset, x, y int, width, height Sizei)
var CopyTexSubImage3D func(target Enum, level, xoffset, yoffset, zoffset int, x, y int, width, height Sizei)
var CopyTexSubImage3DEXT func(target Enum, level, xoffset, yoffset, zoffset, x, y int, width, height Sizei)

// var CopyTextureImage1DEXT func(texture uint, target Enum, level int, internalformat Enum, x, y int, width Sizei, border int)
// var CopyTextureImage2DEXT func(texture uint, target Enum, level int, internalformat Enum, x, y int, width, height Sizei, border int)
// var CopyTextureSubImage1DEXT func(texture uint, target Enum, level, xoffset, x, y int, width Sizei)
// var CopyTextureSubImage2DEXT func(texture uint, target Enum, level, xoffset, yoffset, x, y int, width, height Sizei)
// var CopyTextureSubImage3DEXT func(texture uint, target Enum, level, xoffset, yoffset, zoffset, x, y int, width, height Sizei)
// var CreateDebugObjectMESA func() HandleARB
var CreateProgram func() uint
var CreateProgramObjectARB func() HandleARB
var CreateShader func(type_ Enum) uint
var CreateShaderObjectARB func(shaderType Enum) HandleARB
var CreateShaderProgramEXT func(type_ Enum, str string) uint

// var CreateShaderProgramv func(type_ Enum, count Sizei, strings []string) uint
// var CreateSyncFromCLeventARB func(context *_cl_context, event *_cl_event, flags GLbitfield) GLsync
var CullFace func(mode Enum)

// var CullParameterdvEXT func(pname Enum, params *float64)
// var CullParameterfvEXT func(pname Enum, params *float32)
// var CurrentPaletteMatrixARB func(index int)
// var DebugMessageCallbackAMD func(callback GLDEBUGPROCAMD, userParam *GLvoid)
// var DebugMessageCallbackARB func(callback GLDEBUGPROCARB, userParam *GLvoid)
// var DebugMessageControlARB func(source, type_, severity Enum, count Sizei, ids *uint, enabled GLboolean)
// var DebugMessageEnableAMD func(category, severity Enum, count Sizei, ids *uint, enabled GLboolean)
// var DebugMessageInsertAMD func(category, severity Enum, id uint, length Sizei, buf string)
// var DebugMessageInsertARB func(source, type_ Enum, id uint, severity Enum, length Sizei, buf string)
// var DeformationMap3dSGIX func(target Enum, u1, u2 float64, ustride, uorder int, v1, v2 float64, vstride, vorder int, w1, w2 float64, wstride, worder int, points *float64)
// var DeformationMap3fSGIX func(target Enum, u1, u2 float32, ustride, uorder int, v1, v2 float32, vstride, vorder int, w1, w2 float32, wstride, worder int, points *float32)
// var DeformSGIX func(mask GLbitfield)
// var DeleteAsyncMarkersSGIX func(marker uint, range_ Sizei)
var DeleteBuffers func(n Sizei, buffers *uint)
var DeleteBuffersARB func(n Sizei, buffers *uint)

// var DeleteFencesAPPLE func(n Sizei, fences *uint)
// var DeleteFencesNV func(n Sizei, fences *uint)
var DeleteFragmentShaderATI func(id uint)
var DeleteFramebuffers func(n Sizei, framebuffers *uint)
var DeleteFramebuffersEXT func(n Sizei, framebuffers *uint)
var DeleteLists func(list uint, range_ Sizei)

// var DeleteNamedStringARB func(namelen int, name string)
// var DeleteNamesAMD func(identifier Enum, num uint, names *uint)
var DeleteObjectARB func(obj HandleARB)

// var DeleteOcclusionQueriesNV func(n Sizei, ids *uint)
// var DeletePerfMonitorsAMD func(n Sizei, monitors *uint)
var DeleteProgram func(program uint)

// var DeleteProgramPipelines func(n Sizei, pipelines *uint)
var DeleteProgramsARB func(n Sizei, programs *uint)
var DeleteProgramsNV func(n Sizei, programs *uint)
var DeleteQueries func(n Sizei, ids *uint)
var DeleteQueriesARB func(n Sizei, ids *uint)
var DeleteRenderbuffers func(n Sizei, renderbuffers *uint)
var DeleteRenderbuffersEXT func(n Sizei, renderbuffers *uint)
var DeleteSamplers func(count Sizei, samplers *uint)
var DeleteShader func(shader uint)
var DeleteSync func(sync Sync)
var DeleteTextures func(n Sizei, textures *uint)
var DeleteTexturesEXT func(n Sizei, textures *uint)
var DeleteTransformFeedbacks func(n Sizei, ids *uint)

// var DeleteTransformFeedbacksNV func(n Sizei, ids *uint)
var DeleteVertexArrays func(n Sizei, arrays *uint)

// var DeleteVertexArraysAPPLE func(n Sizei, arrays *uint)
// var DeleteVertexShaderEXT func(id uint)
// var DepthBoundsdNV func(zmin, zmax float64)
// var DepthBoundsEXT func(zmin, zmax clampd)
var DepthFunc func(func_ Enum)
var DepthMask func(flag GLboolean)
var DepthRange func(near_val, far_val Clampd)

// var DepthRangeArrayv func(first uint, count Sizei, v *clampd)
// var DepthRangedNV func(zNear, zFar float64)
var DepthRangef func(n, f Clampf)

// var DepthRangeIndexed func(index uint, n, f clampd)
var DetachObjectARB func(containerObj, attachedObj HandleARB)
var DetachShader func(program, shader uint)

// var DetailTexFuncSGIS func(target Enum, n Sizei, points *float32)
var Disable func(cap Enum)
var DisableClientState func(cap Enum)

// var DisableClientStateIndexedEXT func(array Enum, index uint)
var Disablei func(target Enum, index uint)
var DisableIndexedEXT func(target Enum, index uint)

// var DisableVariantClientStateEXT func(id uint)
// var DisableVertexAttribAPPLE func(index uint, pname Enum)
var DisableVertexAttribArray func(index uint)
var DisableVertexAttribArrayARB func(index uint)
var DrawArrays func(mode Enum, first int, count Sizei)
var DrawArraysEXT func(mode Enum, first int, count Sizei)

// var DrawArraysIndirect func(mode Enum, indirect *GLvoid)
var DrawArraysInstanced func(mode Enum, first int, count, primcount Sizei)
var DrawArraysInstancedARB func(mode Enum, first int, count, primcount Sizei)

// var DrawArraysInstancedBaseInstance func(mode Enum, first int, count, primcount Sizei, baseinstance uint)
var DrawArraysInstancedEXT func(mode Enum, start int, count, primcount Sizei)
var DrawBuffer func(mode Enum)
var DrawBuffers func(n Sizei, bufs *Enum)
var DrawBuffersARB func(n Sizei, bufs *Enum)
var DrawBuffersATI func(n Sizei, bufs *Enum)

// var DrawElementArrayAPPLE func(mode Enum, first int, count Sizei)
// var DrawElementArrayATI func(mode Enum, count Sizei)
var DrawElements func(mode Enum, count Sizei, type_ Enum, indices *GLvoid)
var DrawElementsBaseVertex func(mode Enum, count Sizei, type_ Enum, indices *GLvoid, basevertex int)

// var DrawElementsIndirect func(mode, type_ Enum, indirect *GLvoid)
var DrawElementsInstanced func(mode Enum, count Sizei, type_ Enum, indices *GLvoid, primcount Sizei)
var DrawElementsInstancedARB func(mode Enum, count Sizei, type_ Enum, indices *GLvoid, primcount Sizei)

// var DrawElementsInstancedBaseInstance func(mode Enum, count Sizei, type_ Enum, indices *void, primcount Sizei, baseinstance uint)
var DrawElementsInstancedBaseVertex func(mode Enum, count Sizei, type_ Enum, indices *GLvoid, primcount Sizei, basevertex int)

// var DrawElementsInstancedBaseVertexBaseInstance func(mode Enum, count Sizei, type_ Enum, indices *void, primcount Sizei, basevertex int, baseinstance uint)
var DrawElementsInstancedEXT func(mode Enum, count Sizei, type_ Enum, indices *GLvoid, primcount Sizei)

// var DrawMeshArraysSUN func(mode Enum, first int, count, width Sizei)
var DrawPixels func(width, height Sizei, format, type_ Enum, pixels *GLvoid)

// var DrawRangeElementArrayAPPLE func(mode Enum, start, end uint, first int, count Sizei)
// var DrawRangeElementArrayATI func(mode Enum, start, end uint, count Sizei)
var DrawRangeElements func(mode Enum, start, end uint, count Sizei, type_ Enum, indices *GLvoid)
var DrawRangeElementsBaseVertex func(mode Enum, start, end uint, count Sizei, type_ Enum, indices *GLvoid, basevertex int)
var DrawRangeElementsEXT func(mode Enum, start, end uint, count Sizei, type_ Enum, indices *GLvoid)
var DrawTransformFeedback func(mode Enum, id uint)

// var DrawTransformFeedbackInstanced func(mode Enum, id uint, primcount Sizei)
// var DrawTransformFeedbackNV func(mode Enum, id uint)
// var DrawTransformFeedbackStream func(mode Enum, id, stream uint)
// var DrawTransformFeedbackStreamInstanced func(mode Enum, id, stream uint, primcount Sizei)
var EdgeFlag func(flag GLboolean)

// var EdgeFlagFormatNV func(stride Sizei)
var EdgeFlagPointer func(stride Sizei, ptr *GLvoid)
var EdgeFlagPointerEXT func(stride, count Sizei, pointer *GLboolean)

// var EdgeFlagPointerListIBM func(stride int, pointer **GLboolean, ptrstride int)
var EdgeFlagv func(flag *GLboolean)
var EGLImageTargetRenderbufferStorageOES func(target Enum, image EglImageOES)
var EGLImageTargetTexture2DOES func(target Enum, image EglImageOES)

// var ElementPointerAPPLE func(type_ Enum, pointer *GLvoid)
// var ElementPointerATI func(type_ Enum, pointer *GLvoid)
var Enable func(cap Enum)
var EnableClientState func(cap Enum)

// var EnableClientStateIndexedEXT func(array Enum, index uint)
var Enablei func(target Enum, index uint)
var EnableIndexedEXT func(target Enum, index uint)

// var EnableVariantClientStateEXT func(id uint)
// var EnableVertexAttribAPPLE func(index uint, pname Enum)
var EnableVertexAttribArray func(index uint)
var EnableVertexAttribArrayARB func(index uint)
var End func()
var EndConditionalRender func()
var EndConditionalRenderNV func()
var EndFragmentShaderATI func()
var EndList func()

// var EndOcclusionQueryNV func()
// var EndPerfMonitorAMD func(monitor uint)
var EndQuery func(target Enum)
var EndQueryARB func(target Enum)

// var EndQueryIndexed func(target Enum, index uint)
var EndTransformFeedback func()
var EndTransformFeedbackEXT func()

// var EndTransformFeedbackNV func()
// var EndVertexShaderEXT func()
// var EndVideoCaptureNV func(video_capture_slot uint)
var EvalCoord1d func(u float64)
var EvalCoord1dv func(u *float64)
var EvalCoord1f func(u float32)
var EvalCoord1fv func(u *float32)
var EvalCoord2d func(u, v float64)
var EvalCoord2dv func(u *float64)
var EvalCoord2f func(u, v float32)
var EvalCoord2fv func(u *float32)

// var EvalMapsNV func(target, mode Enum)
var EvalMesh1 func(mode Enum, i1, i2 int)
var EvalMesh2 func(mode Enum, hi1, i2, j1, j2 int)
var EvalPoint1 func(i int)
var EvalPoint2 func(i, j int)
var ExecuteProgramNV func(target Enum, id uint, params *float32)

// var ExtractComponentEXT func(res, src, num uint)
var FeedbackBuffer func(size Sizei, type_ Enum, buffer *float32)
var FenceSync func(condition Enum, flags GLbitfield) Sync
var FinalCombinerInputNV func(variable, input, mapping, componentUsage Enum)
var Finish func()

// var FinishAsyncSGIX func(markerp *uint) int
// var FinishFenceAPPLE func(fence uint)
// var FinishFenceNV func(fence uint)
// var FinishObjectAPPLE func(object Enum, name int)
// var FinishTextureSUNX func()
var Flush func()
var FlushMappedBufferRange func(target Enum, offset Intptr, length Sizeiptr)

// var FlushMappedBufferRangeAPPLE func(target Enum, offset Intptr, size Sizeiptr)
// var FlushMappedNamedBufferRangeEXT func(buffer uint, offset Intptr, length Sizeiptr)
// var FlushPixelDataRangeNV func(target Enum)
// var FlushRasterSGIX func()
// var FlushVertexArrayRangeAPPLE func(length Sizei, pointer *GLvoid)
var FlushVertexArrayRangeNV func()
var FogCoordd func(coord float64)
var FogCoorddEXT func(coord float64)
var FogCoorddv func(coord *float64)
var FogCoorddvEXT func(coord *float64)
var FogCoordf func(coord float32)
var FogCoordfEXT func(coord float32)

// var FogCoordFormatNV func(type_ Enum, stride Sizei)
var FogCoordfv func(coord *float32)
var FogCoordfvEXT func(coord *float32)

// var FogCoordhNV func(fog GLhalfNV)
// var FogCoordhvNV func(fog *GLhalfNV)
var FogCoordPointer func(type_ Enum, stride Sizei, pointer *GLvoid)
var FogCoordPointerEXT func(type_ Enum, stride Sizei, pointer *GLvoid)

// var FogCoordPointerListIBM func(type_ Enum, stride int, pointer **GLvoid, ptrstride int)
var Fogf func(pname Enum, param float32)

// var FogFuncSGIS func(n Sizei, points *float32)
var Fogfv func(pname Enum, params *float32)
var Fogi func(pname Enum, param int)
var Fogiv func(pname Enum, params *int)

// var FragmentColorMaterialSGIX func(face, mode Enum)
// var FragmentLightfSGIX func(light, pname Enum, param float32)
// var FragmentLightfvSGIX func(light, pname Enum, params *float32)
// var FragmentLightiSGIX func(light, pname Enum, param int)
// var FragmentLightivSGIX func(light, pname Enum, params *int)
// var FragmentLightModelfSGIX func(pname Enum, param float32)
// var FragmentLightModelfvSGIX func(pname Enum, params *float32)
// var FragmentLightModeliSGIX func(pname Enum, param int)
// var FragmentLightModelivSGIX func(pname Enum, params *int)
// var FragmentMaterialfSGIX func(face, pname Enum, param float32)
// var FragmentMaterialfvSGIX func(face, pname Enum, params *float32)
// var FragmentMaterialiSGIX func(face, pname Enum, param int)
// var FragmentMaterialivSGIX func(face, pname Enum, params *int)
// var FramebufferDrawBufferEXT func(framebuffer uint, mode Enum)
// var FramebufferDrawBuffersEXT func(framebuffer uint, n Sizei, bufs *Enum)
// var FramebufferReadBufferEXT func(framebuffer uint, mode Enum)
var FramebufferRenderbuffer func(target, attachment, renderbuffertarget Enum, renderbuffer uint)
var FramebufferRenderbufferEXT func(target, attachment, renderbuffertarget Enum, renderbuffer uint)
var FramebufferTexture func(target, attachment Enum, texture uint, level int)
var FramebufferTexture1D func(target, attachment, textarget Enum, texture uint, level int)
var FramebufferTexture1DEXT func(target, attachment, textarget Enum, texture uint, level int)
var FramebufferTexture2D func(target, attachment, textarget Enum, texture uint, level int)
var FramebufferTexture2DEXT func(target, attachment, textarget Enum, texture uint, level int)
var FramebufferTexture3D func(target, attachment, textarget Enum, texture uint, level, zoffset int)
var FramebufferTexture3DEXT func(target, attachment, textarget Enum, texture uint, level, zoffset int)
var FramebufferTextureARB func(target, attachment Enum, texture uint, level int)

// var FramebufferTextureEXT func(target, attachment Enum, texture uint, level int)
var FramebufferTextureFaceARB func(target, attachment Enum, texture uint, level int, face Enum)

// var FramebufferTextureFaceEXT func(target, attachment Enum, texture uint, level int, face Enum)
var FramebufferTextureLayer func(target, attachment Enum, texture uint, level, layer int)
var FramebufferTextureLayerARB func(target, attachment Enum, texture uint, level, layer int)
var FramebufferTextureLayerEXT func(target, attachment Enum, texture uint, level, layer int)

// var FrameTerminatorGREMEDY func()
// var FrameZoomSGIX func(factor int)
// var FreeObjectBufferATI func(buffer uint)
var FrontFace func(mode Enum)
var Frustum func(hhleft, right, bottom, top, near_val, far_val float64)

// var GenAsyncMarkersSGIX func(range_ Sizei) uint
var GenBuffers func(n Sizei, buffers *uint)
var GenBuffersARB func(n Sizei, buffers *uint)
var GenerateMipmap func(target Enum)
var GenerateMipmapEXT func(target Enum)

// var GenerateMultiTexMipmapEXT func(texunit, target Enum)
// var GenerateTextureMipmapEXT func(texture uint, target Enum)
// var GenFencesAPPLE func(n Sizei, fences *uint)
// var GenFencesNV func(n Sizei, fences *uint)
var GenFragmentShadersATI func(range_ uint) uint
var GenFramebuffers func(n Sizei, framebuffers *uint)
var GenFramebuffersEXT func(n Sizei, framebuffers *uint)
var GenLists func(range_ Sizei) uint

// var GenNamesAMD func(identifier Enum, num uint, names *uint)
// var GenOcclusionQueriesNV func(n Sizei, ids *uint)
// var GenPerfMonitorsAMD func(n Sizei, monitors *uint)
// var GenProgramPipelines func(n Sizei, pipelines *uint)
var GenProgramsARB func(n Sizei, programs *uint)
var GenProgramsNV func(n Sizei, programs *uint)
var GenQueries func(n Sizei, ids *uint)
var GenQueriesARB func(n Sizei, ids *uint)
var GenRenderbuffers func(n Sizei, renderbuffers *uint)
var GenRenderbuffersEXT func(n Sizei, renderbuffers *uint)
var GenSamplers func(count Sizei, samplers *uint)

// var GenSymbolsEXT func(datatype, storagetype, range_ Enum, components uint) uint
var GenTextures func(n Sizei, textures *uint)
var GenTexturesEXT func(n Sizei, textures *uint)
var GenTransformFeedbacks func(n Sizei, ids *uint)

// var GenTransformFeedbacksNV func(n Sizei, ids *uint)
var GenVertexArrays func(n Sizei, arrays *uint)

// var GenVertexArraysAPPLE func(n Sizei, arrays *uint)
// var GenVertexShadersEXT func(range_ uint) uint
// var GetActiveAtomicCounterBufferiv func(program, bufferIndex uint, pname Enum, params *int)
var GetActiveAttrib func(program, index uint, bufSize Sizei, length *Sizei, size *int, type_ *Enum, name string)
var GetActiveAttribARB func(programObj HandleARB, index uint, maxLength Sizei, length *Sizei, size *int, type_ *Enum, name *GLcharARB)

// var GetActiveSubroutineName func(program uint, shadertype_ Enum, index uint, bufsize Sizei, length *Sizei, name string)
// var GetActiveSubroutineUniformiv func(program uint, shadertype Enum, index uint, pname Enum, values *int)
// var GetActiveSubroutineUniformName func(program uint, shadertype Enum, index uint, bufsize Sizei, length *Sizei, name string)
var GetActiveUniform func(program, index uint, bufSize Sizei, length *Sizei, size *int, type_ *Enum, name string)
var GetActiveUniformARB func(programObj HandleARB, index uint, maxLength Sizei, length *Sizei, size *int, type_ *Enum, name *GLcharARB)

// var GetActiveUniformBlockiv func(program, uniformBlockIndex uint, pname Enum, params *int)
// var GetActiveUniformBlockName func(program, uniformBlockIndex uint, bufSize Sizei, length *Sizei, uniformBlockName string)
// var GetActiveUniformName func(program, uniformIndex uint, bufSize Sizei, length *Sizei, uniformName string)
// var GetActiveUniformsiv func(program uint, uniformCount Sizei, uniformIndices *uint, pname Enum, params *int)
// var GetActiveVaryingNV func(program, index uint, bufSize Sizei, length, size *Sizei, type_ *Enum, name string)
// var GetArrayObjectfvATI func(array, pname Enum, params *float32)
// var GetArrayObjectivATI func(array, pname Enum, params *int)
var GetAttachedObjectsARB func(containerObj HandleARB, maxCount Sizei, count *Sizei, obj *HandleARB)
var GetAttachedShaders func(program uint, maxCount Sizei, count *Sizei, obj *uint)
var GetAttribLocation func(program uint, name string) int
var GetAttribLocationARB func(programObj HandleARB, name *GLcharARB) int
var GetBooleani_v func(target Enum, index uint, data *GLboolean)
var GetBooleanIndexedvEXT func(target Enum, index uint, data *GLboolean)
var GetBooleanv func(pname Enum, params *GLboolean)
var GetBufferParameteri64v func(target, pname Enum, params *int64)
var GetBufferParameteriv func(target, pname Enum, params *int)
var GetBufferParameterivARB func(target, pname Enum, params *int)

// var GetBufferParameterui64vNV func(target, pname Enum, params *uint64EXT)
var GetBufferPointerv func(target, pname Enum, params **GLvoid)
var GetBufferPointervARB func(target, pname Enum, params **GLvoid)
var GetBufferSubData func(target Enum, offset Intptr, size Sizeiptr, data *GLvoid)
var GetBufferSubDataARB func(target Enum, offset IntptrARB, size SizeiptrARB, data *GLvoid)
var GetClipPlane func(plane Enum, equation *float64)
var GetColorTable func(target, format, type_ Enum, table *GLvoid)
var GetColorTableEXT func(target, format, type_ Enum, data *GLvoid)
var GetColorTableParameterfv func(target, pname Enum, params *float32)
var GetColorTableParameterfvEXT func(target, pname Enum, params *float32)

// var GetColorTableParameterfvSGI func(target, pname Enum, params *float32)
var GetColorTableParameteriv func(target, pname Enum, params *int)
var GetColorTableParameterivEXT func(target, pname Enum, params *int)

// var GetColorTableParameterivSGI func(target, pname Enum, params *int)
// var GetColorTableSGI func(target, format, type_ Enum, table *GLvoid)
var GetCombinerInputParameterfvNV func(stage, portion, variable, pname Enum, params *float32)
var GetCombinerInputParameterivNV func(stage, portion, variable, pname Enum, params *int)
var GetCombinerOutputParameterfvNV func(stage, portion, pname Enum, params *float32)
var GetCombinerOutputParameterivNV func(stage, portion, pname Enum, params *int)

// var GetCombinerStageParameterfvNV func(stage, pname Enum, params *float32)
// var GetCompressedMultiTexImageEXT func(texunit, target Enum, lod int, img *GLvoid)
var GetCompressedTexImage func(target Enum, lod int, img *GLvoid)
var GetCompressedTexImageARB func(target Enum, level int, img *GLvoid)

// var GetCompressedTextureImageEXT func(texture uint, target Enum, lod int, img *GLvoid)
var GetConvolutionFilter func(target, format, type_ Enum, image *GLvoid)

// var GetConvolutionFilterEXT func(target, format, type_ Enum, image *GLvoid)
var GetConvolutionParameterfv func(target, pname Enum, params *float32)

// var GetConvolutionParameterfvEXT func(target, pname Enum, params *float32)
var GetConvolutionParameteriv func(target, pname Enum, params *int)

// var GetConvolutionParameterivEXT func(target, pname Enum, params *int)
// var GetDebugLogLengthMESA func(obj HandleARB, logType, shaderType Enum) Sizei
// var GetDebugLogMESA func(obj HandleARB, logType, shaderType Enum, maxLength Sizei, length *Sizei, debugLog *GLcharARB)
// var GetDebugMessageLogAMD func(count uint, bufsize Sizei, categories *Enum, severities, ids *uint, lengths *Sizei, message string) uint
// var GetDebugMessageLogARB func(count uint, bufsize Sizei, sources, types *Enum, ids *uint, severities *Enum, lengths *Sizei, messageLog string) uint
// var GetDetailTexFuncSGIS func(target Enum, points *float32)
// var GetDoublei_v func(target Enum, index uint, data *float64)
// var GetDoubleIndexedvEXT func(target Enum, index uint, data *float64)
var GetDoublev func(pname Enum, params *float64)
var GetError func() Enum

// var GetFenceivNV func(fence uint, pname Enum, params *int)
var GetFinalCombinerInputParameterfvNV func(variable, pname Enum, params *float32)
var GetFinalCombinerInputParameterivNV func(variable, pname Enum, params *int)

// var GetFloati_v func(target Enum, index uint, data *float32)
// var GetFloatIndexedvEXT func(target Enum, index uint, data *float32)
var GetFloatv func(pname Enum, params *float32)

// var GetFogFuncSGIS func(points *float32)
// var GetFragDataIndex func(program uint, name string) int
var GetFragDataLocation func(program uint, name string) int
var GetFragDataLocationEXT func(program uint, name string) int

// var GetFragmentLightfvSGIX func(light, pname Enum, params *float32)
// var GetFragmentLightivSGIX func(light, pname Enum, params *int)
// var GetFragmentMaterialfvSGIX func(face, pname Enum, params *float32)
// var GetFragmentMaterialivSGIX func(face, pname Enum, params *int)
var GetFramebufferAttachmentParameteriv func(target, attachment, pname Enum, params *int)
var GetFramebufferAttachmentParameterivEXT func(target, attachment, pname Enum, params *int)

// var GetFramebufferParameterivEXT func(framebuffer uint, pname Enum, params *int)
var GetGraphicsResetStatusARB func() Enum
var GetHandleARB func(pname Enum) HandleARB
var GetHistogram func(target Enum, reset GLboolean, format, type_ Enum, values *GLvoid)

// var GetHistogramEXT func(target Enum, reset GLboolean, format, type_ Enum, values *GLvoid)
var GetHistogramParameterfv func(target, pname Enum, params *float32)

// var GetHistogramParameterfvEXT func(target, pname Enum, params *float32)
var GetHistogramParameteriv func(target, pname Enum, params *int)

// var GetHistogramParameterivEXT func(target, pname Enum, params *int)
// var GetImageTransformParameterfvHP func(target, pname Enum, params *float32)
// var GetImageTransformParameterivHP func(target, pname Enum, params *int)
var GetInfoLogARB func(obj HandleARB, maxLength Sizei, length *Sizei, infoLog *GLcharARB)

// var GetInstrumentsSGIX func() int
var GetInteger64i_v func(target Enum, index uint, data *int64)
var GetInteger64v func(pname Enum, params *int64)
var GetIntegeri_v func(target Enum, index uint, data *int)
var GetIntegerIndexedvEXT func(target Enum, index uint, data *int)

// var GetIntegerui64i_vNV func(value Enum, index uint, result *uint64EXT)
// var GetIntegerui64vNV func(value Enum, result *uint64EXT)
var GetIntegerv func(pname Enum, params *int)

// var GetInternalformativ func(target, internalformat, pname Enum, bufSize Sizei, params *int)
// var GetInvariantBooleanvEXT func(id uint, value Enum, data *GLboolean)
// var GetInvariantFloatvEXT func(id uint, value Enum, data *float32)
// var GetInvariantIntegervEXT func(id uint, value Enum, data *int)
var GetLightfv func(light, pname Enum, params *float32)
var GetLightiv func(light, pname Enum, params *int)

// var GetListParameterfvSGIX func(list uint, pname Enum, params *float32)
// var GetListParameterivSGIX func(list uint, pname Enum, params *int)
// var GetLocalConstantBooleanvEXT func(id uint, value Enum, data *GLboolean)
// var GetLocalConstantFloatvEXT func(id uint, value Enum, data *float32)
// var GetLocalConstantIntegervEXT func(id uint, value Enum, data *int)
// var GetMapAttribParameterfvNV func(target Enum, index uint, pname Enum, params *float32)
// var GetMapAttribParameterivNV func(target Enum, index uint, pname Enum, params *int)
// var GetMapControlPointsNV func(target Enum, index uint, type_ Enum, ustride, vstride Sizei, packed GLboolean, points *GLvoid)
var GetMapdv func(target, query Enum, v *float64)
var GetMapfv func(target, query Enum, v *float32)
var GetMapiv func(target, query Enum, v *int)

// var GetMapParameterfvNV func(target, pname Enum, params *float32)
// var GetMapParameterivNV func(target, pname Enum, params *int)
var GetMaterialfv func(face, pname Enum, params *float32)
var GetMaterialiv func(face, pname Enum, params *int)
var GetMinmax func(target Enum, reset GLboolean, format, types Enum, values *GLvoid)

// var GetMinmaxEXT func(target Enum, reset GLboolean, format, type_ Enum, values *GLvoid)
var GetMinmaxParameterfv func(target, pname Enum, params *float32)

// var GetMinmaxParameterfvEXT func(target, pname Enum, params *float32)
var GetMinmaxParameteriv func(target, pname Enum, params *int)

// var GetMinmaxParameterivEXT func(target, pname Enum, params *int)
// var GetMultisamplefv func(pname Enum, index uint, val *float32)
// var GetMultisamplefvNV func(pname Enum, index uint, val *float32)
// var GetMultiTexEnvfvEXT func(texunit, target, pname Enum, params *float32)
// var GetMultiTexEnvivEXT func(texunit, target, pname Enum, params *int)
// var GetMultiTexGendvEXT func(texunit, coord, pname Enum, params *float64)
// var GetMultiTexGenfvEXT func(texunit, coord, pname Enum, params *float32)
// var GetMultiTexGenivEXT func(texunit, coord, pname Enum, params *int)
// var GetMultiTexImageEXT func(texunit, target Enum, level int, format, type_ Enum, pixels *GLvoid)
// var GetMultiTexLevelParameterfvEXT func(texunit, target Enum, level int, pname Enum, params *float32)
// var GetMultiTexLevelParameterivEXT func(texunit, target Enum, level int, pname Enum, params *int)
// var GetMultiTexParameterfvEXT func(texunit, target, pname Enum, params *float32)
// var GetMultiTexParameterIivEXT func(texunit, target, pname Enum, params *int)
// var GetMultiTexParameterIuivEXT func(texunit, target, pname Enum, params *uint)
// var GetMultiTexParameterivEXT func(texunit, target, pname Enum, params *int)
// var GetNamedBufferParameterivEXT func(buffer uint, pname Enum, params *int)
// var GetNamedBufferParameterui64vNV func(buffer uint, pname Enum, params *uint64EXT)
// var GetNamedBufferPointervEXT func(buffer uint, pname Enum, params **GLvoid)
// var GetNamedBufferSubDataEXT func(buffer uint, offset Intptr, size Sizeiptr, data *GLvoid)
// var GetNamedFramebufferAttachmentParameterivEXT func(framebuffer uint, attachment, pname Enum, params *int)
// var GetNamedProgramivEXT func(program uint, target, pname Enum, params *int)
// var GetNamedProgramLocalParameterdvEXT func(program uint, target Enum, index uint, params *float64)
// var GetNamedProgramLocalParameterfvEXT func(program uint, target Enum, index uint, params *float32)
// var GetNamedProgramLocalParameterIivEXT func(program uint, target Enum, index uint, params *int)
// var GetNamedProgramLocalParameterIuivEXT func(program uint, target Enum, index uint, params *uint)
// var GetNamedProgramStringEXT func(program uint, target, pname Enum, string *GLvoid)
// var GetNamedRenderbufferParameterivEXT func(renderbuffer uint, pname Enum, params *int)
// var GetNamedStringARB func(namelen int, name string, bufSize Sizei, stringlen *int, string string)
// var GetNamedStringivARB func(namelen int, name string, pname Enum, params *int)
var GetnColorTableARB func(target, format, type_ Enum, bufSize Sizei, table *GLvoid)
var GetnCompressedTexImageARB func(target Enum, lod int, bufSize Sizei, img *GLvoid)
var GetnConvolutionFilterARB func(target, format, type_ Enum, bufSize Sizei, image *GLvoid)
var GetnHistogramARB func(target Enum, reset GLboolean, format, type_ Enum, bufSize Sizei, values *GLvoid)
var GetnMapdvARB func(target, query Enum, bufSize Sizei, v *float64)
var GetnMapfvARB func(target, query Enum, bufSize Sizei, v *float32)
var GetnMapivARB func(target, query Enum, bufSize Sizei, v *int)
var GetnMinmaxARB func(target Enum, reset GLboolean, format, type_ Enum, bufSize Sizei, values *GLvoid)
var GetnPixelMapfvARB func(map_ Enum, bufSize Sizei, values *float32)
var GetnPixelMapuivARB func(map_ Enum, bufSize Sizei, values *uint)
var GetnPixelMapusvARB func(map_ Enum, bufSize Sizei, values *uint16)
var GetnPolygonStippleARB func(bufSize Sizei, pattern *uint8)
var GetnSeparableFilterARB func(target, format, type_ Enum, rowBufSize Sizei, row *GLvoid, columnBufSize Sizei, column, span *GLvoid)
var GetnTexImageARB func(target Enum, level int, format, type_ Enum, bufSize Sizei, img *GLvoid)
var GetnUniformdvARB func(program uint, location int, bufSize Sizei, params *float64)
var GetnUniformfvARB func(program uint, location int, bufSize Sizei, params *float32)
var GetnUniformivARB func(program uint, location int, bufSize Sizei, params *int)
var GetnUniformuivARB func(program uint, location int, bufSize Sizei, params *uint)

// var GetObjectBufferfvATI func(buffer uint, pname Enum, params *float32)
// var GetObjectBufferivATI func(buffer uint, pname Enum, params *int)
var GetObjectParameterfvARB func(obj HandleARB, pname Enum, params *float32)
var GetObjectParameterivAPPLE func(objectType Enum, name uint, pname Enum, params *int)
var GetObjectParameterivARB func(obj HandleARB, pname Enum, params *int)

// var GetOcclusionQueryivNV func(id uint, pname Enum, params *int)
// var GetOcclusionQueryuivNV func(id uint, pname Enum, params *uint)
// var GetPerfMonitorCounterDataAMD func(monitor uint, pname Enum, dataSize Sizei, data *uint, bytesWritten *int)
// var GetPerfMonitorCounterInfoAMD func(group, counter uint, pname Enum, data *GLvoid)
// var GetPerfMonitorCountersAMD func(group uint, numCounters, maxActiveCounters *int, counterSize Sizei, counters *uint)
// var GetPerfMonitorCounterStringAMD func(group, counter uint, bufSize Sizei, length *Sizei, counterString string)
// var GetPerfMonitorGroupsAMD func(numGroups *int, groupsSize Sizei, groups *uint)
// var GetPerfMonitorGroupStringAMD func(group uint, bufSize Sizei, length *Sizei, groupString string)
var GetPixelMapfv func(map_ Enum, values *float32)
var GetPixelMapuiv func(map_ Enum, values *uint)
var GetPixelMapusv func(map_ Enum, values *uint16)

// var GetPixelTexGenParameterfvSGIS func(pname Enum, params *float32)
// var GetPixelTexGenParameterivSGIS func(pname Enum, params *int)
// var GetPointerIndexedvEXT func(target Enum, index uint, data **GLvoid)
var GetPointerv func(pname Enum, params **GLvoid)
var GetPointervEXT func(pname Enum, params **GLvoid)
var GetPolygonStipple func(mask *uint8)

// var GetProgramBinary func(program uint, bufSize Sizei, length *Sizei, binaryFormat *Enum, binary *GLvoid)
var GetProgramEnvParameterdvARB func(target Enum, index uint, params *float64)
var GetProgramEnvParameterfvARB func(target Enum, index uint, params *float32)

// var GetProgramEnvParameterIivNV func(target Enum, index uint, params *int)
// var GetProgramEnvParameterIuivNV func(target Enum, index uint, params *uint)
var GetProgramInfoLog func(program uint, bufSize Sizei, length *Sizei, infoLog string)
var GetProgramiv func(program uint, pname Enum, params *int)
var GetProgramivARB func(target, pname Enum, params *int)
var GetProgramivNV func(id uint, pname Enum, params *int)
var GetProgramLocalParameterdvARB func(target Enum, index uint, params *float64)
var GetProgramLocalParameterfvARB func(target Enum, index uint, params *float32)

// var GetProgramLocalParameterIivNV func(target Enum, index uint, params *int)
// var GetProgramLocalParameterIuivNV func(target Enum, index uint, params *uint)
var GetProgramNamedParameterdvNV func(id uint, len Sizei, name *uint8, params *float64)
var GetProgramNamedParameterfvNV func(id uint, len Sizei, name *uint8, params *float32)
var GetProgramParameterdvNV func(target Enum, index uint, pname Enum, params *float64)
var GetProgramParameterfvNV func(target Enum, index uint, pname Enum, params *float32)

// var GetProgramPipelineInfoLog func(pipeline uint, bufSize Sizei, length *Sizei, infoLog string)
// var GetProgramPipelineiv func(pipeline uint, pname Enum, params *int)
// var GetProgramRegisterfvMESA func(target Enum, len Sizei, name *uint8, v *float32)
// var GetProgramStageiv func(program uint, shadertype_, pname Enum, values *int)
var GetProgramStringARB func(target, pname Enum, string *GLvoid)
var GetProgramStringNV func(id uint, pname Enum, program *uint8)

// var GetProgramSubroutineParameteruivNV func(target Enum, index uint, param *uint)
// var GetQueryIndexediv func(target Enum, index uint, pname Enum, params *int)
var GetQueryiv func(target, pname Enum, params *int)
var GetQueryivARB func(target, pname Enum, params *int)

// var GetQueryObjecti64v func(id uint, pname Enum, params *int64)
// var GetQueryObjecti64vEXT func(id uint, pname Enum, params *int64EXT)
var GetQueryObjectiv func(id uint, pname Enum, params *int)
var GetQueryObjectivARB func(id uint, pname Enum, params *int)

// var GetQueryObjectui64v func(id uint, pname Enum, params *uint64)
// var GetQueryObjectui64vEXT func(id uint, pname Enum, params *uint64EXT)
var GetQueryObjectuiv func(id uint, pname Enum, params *uint)
var GetQueryObjectuivARB func(id uint, pname Enum, params *uint)
var GetRenderbufferParameteriv func(target, pname Enum, params *int)
var GetRenderbufferParameterivEXT func(target, pname Enum, params *int)
var GetSamplerParameterfv func(sampler uint, pname Enum, params *float32)
var GetSamplerParameterIiv func(sampler uint, pname Enum, params *int)
var GetSamplerParameterIuiv func(sampler uint, pname Enum, params *uint)
var GetSamplerParameteriv func(sampler uint, pname Enum, params *int)
var GetSeparableFilter func(target, format, type_ Enum, row, column, span *GLvoid)

// var GetSeparableFilterEXT func(target, format, type_ Enum, row, column, span *GLvoid)
var GetShaderInfoLog func(shader uint, bufSize Sizei, length *Sizei, infoLog string)
var GetShaderiv func(shader uint, pname Enum, params *int)
var GetShaderPrecisionFormat func(shadertype_, precisiontype Enum, range_, precision *int)
var GetShaderSource func(shader uint, bufSize Sizei, length *Sizei, source string)
var GetShaderSourceARB func(obj HandleARB, maxLength Sizei, length *Sizei, source *GLcharARB)

// var GetSharpenTexFuncSGIS func(target Enum, points *float32)
var GetString func(name String) string
var GetStringi func(name Enum, index uint) *uint8

// var GetSubroutineIndex func(program uint, shadertype Enum, name string) uint
// var GetSubroutineUniformLocation func(program uint, shadertype Enum, name string) int
var GetSynciv func(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *int)
var GetTexBumpParameterfvATI func(pname Enum, param *float32)
var GetTexBumpParameterivATI func(pname Enum, param *int)
var GetTexEnvfv func(target, pname Enum, params *float32)
var GetTexEnviv func(target, pname Enum, params *int)

// var GetTexFilterFuncSGIS func(target, filter Enum, weights *float32)
var GetTexGendv func(coord, pname Enum, params *float64)
var GetTexGenfv func(coord, pname Enum, params *float32)
var GetTexGeniv func(coord, pname Enum, params *int)
var GetTexImage func(target Enum, level int, format, type_ Enum, pixels *GLvoid)
var GetTexLevelParameterfv func(target Enum, level int, pname Enum, params *float32)
var GetTexLevelParameteriv func(target Enum, level int, pname Enum, params *int)
var GetTexParameterfv func(target, pname Enum, params *float32)
var GetTexParameterIiv func(target, pname Enum, params *int)
var GetTexParameterIivEXT func(target, pname Enum, params *int)
var GetTexParameterIuiv func(target, pname Enum, params *uint)
var GetTexParameterIuivEXT func(target, pname Enum, params *uint)
var GetTexParameteriv func(target, pname Enum, params *int)

// var GetTexParameterPointervAPPLE func(target, pname Enum, params **GLvoid)
// var GetTextureImageEXT func(texture uint, target Enum, level int, format, type_ Enum, pixels *GLvoid)
// var GetTextureLevelParameterfvEXT func(texture uint, target Enum, level int, pname Enum, params *float32)
// var GetTextureLevelParameterivEXT func(texture uint, target Enum, level int, pname Enum, params *int)
// var GetTextureParameterfvEXT func(texture uint, target, pname Enum, params *float32)
// var GetTextureParameterIivEXT func(texture uint, target, pname Enum, params *int)
// var GetTextureParameterIuivEXT func(texture uint, target, pname Enum, params *uint)
// var GetTextureParameterivEXT func(texture uint, target, pname Enum, params *int)
var GetTrackMatrixivNV func(target Enum, address uint, pname Enum, params *int)
var GetTransformFeedbackVarying func(program, index uint, bufSize Sizei, length, size *Sizei, type_ *Enum, name string)
var GetTransformFeedbackVaryingEXT func(program, index uint, bufSize Sizei, length, size *Sizei, type_ *Enum, name string)

// var GetTransformFeedbackVaryingNV func(program, index uint, location *int)
// var GetUniformBlockIndex func(program, uniformBlockName string) uint
// var GetUniformBufferSizeEXT func(program uint, location int) int
// var GetUniformdv func(program uint, location int, params *float64)
var GetUniformfv func(program uint, location int, params *float32)
var GetUniformfvARB func(programObj HandleARB, location int, params *float32)

// var GetUniformi64vNV func(program uint, location int, params *int64EXT)
// var GetUniformIndices func(program uint, uniformCount Sizei, uniformNames []string, uniformIndices *uint)
var GetUniformiv func(program uint, location int, params *int)
var GetUniformivARB func(programObj HandleARB, location int, params *int)
var GetUniformLocation func(program uint, name string) int
var GetUniformLocationARB func(programObj HandleARB, name *GLcharARB) int

// var GetUniformOffsetEXT func(program uint, location int) Intptr
// var GetUniformSubroutineuiv func(shadertype_ Enum, location int, params *uint)
// var GetUniformui64vNV func(program uint, location int, params *uint64EXT)
var GetUniformuiv func(program uint, location int, params *uint)
var GetUniformuivEXT func(program uint, location int, params *uint)

// var GetVariantArrayObjectfvATI func(id uint, pname Enum, params *float32)
// var GetVariantArrayObjectivATI func(id uint, pname Enum, params *int)
// var GetVariantBooleanvEXT func(id uint, value Enum, data *GLboolean)
// var GetVariantFloatvEXT func(id uint, value Enum, data *float32)
// var GetVariantIntegervEXT func(id uint, value Enum, data *int)
// var GetVariantPointervEXT func(id uint, value Enum, data **GLvoid)
// var GetVaryingLocationNV func(program uint, name string) int
// var GetVertexAttribArrayObjectfvATI func(index uint, pname Enum, params *float32)
// var GetVertexAttribArrayObjectivATI func(index uint, pname Enum, params *int)
var GetVertexAttribdv func(index uint, pname Enum, params *float64)
var GetVertexAttribdvARB func(index uint, pname Enum, params *float64)
var GetVertexAttribdvNV func(index uint, pname Enum, params *float64)
var GetVertexAttribfv func(index uint, pname Enum, params *float32)
var GetVertexAttribfvARB func(index uint, pname Enum, params *float32)
var GetVertexAttribfvNV func(index uint, pname Enum, params *float32)
var GetVertexAttribIiv func(index uint, pname Enum, params *int)
var GetVertexAttribIivEXT func(index uint, pname Enum, params *int)
var GetVertexAttribIuiv func(index uint, pname Enum, params *uint)
var GetVertexAttribIuivEXT func(index uint, pname Enum, params *uint)
var GetVertexAttribiv func(index uint, pname Enum, params *int)
var GetVertexAttribivARB func(index uint, pname Enum, params *int)
var GetVertexAttribivNV func(index uint, pname Enum, params *int)

// var GetVertexAttribLdv func(index uint, pname Enum, params *float64)
// var GetVertexAttribLdvEXT func(index uint, pname Enum, params *float64)
// var GetVertexAttribLi64vNV func(index uint, pname Enum, params *int64EXT)
// var GetVertexAttribLui64vNV func(index uint, pname Enum, params *uint64EXT)
var GetVertexAttribPointerv func(index uint, pname Enum, pointer **GLvoid)
var GetVertexAttribPointervARB func(index uint, pname Enum, pointer **GLvoid)
var GetVertexAttribPointervNV func(index uint, pname Enum, pointer **GLvoid)

// var GetVideoCaptureivNV func(video_capture_slot uint, pname Enum, params *int)
// var GetVideoCaptureStreamdvNV func(video_capture_slot, stream uint, pname Enum, params *float64)
// var GetVideoCaptureStreamfvNV func(video_capture_slot, stream uint, pname Enum, params *float32)
// var GetVideoCaptureStreamivNV func(video_capture_slot, stream uint, pname Enum, params *int)
// var GetVideoi64vNV func(video_slot uint, pname Enum, params *int64EXT)
// var GetVideoivNV func(video_slot uint, pname Enum, params *int)
// var GetVideoui64vNV func(video_slot uint, pname Enum, params *uint64EXT)
// var GetVideouivNV func(video_slot uint, pname Enum, params *uint)
// var GlobalAlphaFactorbSUN func(factor GLbyte)
// var GlobalAlphaFactordSUN func(factor float64)
// var GlobalAlphaFactorfSUN func(factor float32)
// var GlobalAlphaFactoriSUN func(factor int)
// var GlobalAlphaFactorsSUN func(factor int16)
// var GlobalAlphaFactorubSUN func(factor uint8)
// var GlobalAlphaFactoruiSUN func(factor uint)
// var GlobalAlphaFactorusSUN func(factor uint16)
var Hint func(target, mode Enum)

// var HintPGI func(target Enum, mode int)
var Histogram func(target Enum, width Sizei, internalformat Enum, sink GLboolean)

// var HistogramEXT func(target Enum, width Sizei, internalformat Enum, sink GLboolean)
// var IglooInterfaceSGIX func(pname Enum, params *GLvoid)
// var ImageTransformParameterfHP func(target, pname Enum, param float32)
// var ImageTransformParameterfvHP func(target, pname Enum, params *float32)
// var ImageTransformParameteriHP func(target, pname Enum, param int)
// var ImageTransformParameterivHP func(target, pname Enum, params *int)
// var ImportSyncEXT func(external_sync_type Enum, external_sync Intptr, flags GLbitfield) Sync
var Indexd func(c float64)
var Indexdv func(c *float64)
var Indexf func(c float32)

// var IndexFormatNV func(type_ Enum, stride Sizei)
// var IndexFuncEXT func(func_ Enum, ref clampf)
var Indexfv func(c *float32)
var Indexi func(c int)
var Indexiv func(c *int)
var IndexMask func(mask uint)

// var IndexMaterialEXT func(face, mode Enum)
var IndexPointer func(type_ Enum, stride Sizei, ptr *GLvoid)
var IndexPointerEXT func(type_ Enum, stride, count Sizei, pointer *GLvoid)

// var IndexPointerListIBM func(type_ Enum, stride int, pointer **GLvoid, ptrstride int)
var Indexs func(c int16)
var Indexsv func(c *int16)
var Indexub func(c uint8)
var Indexubv func(c *uint8)
var InitNames func()

// var InsertComponentEXT func(res, src, num uint)
// var InstrumentsBufferSGIX func(size Sizei, buffer *int)
var InterleavedArrays func(format Enum, stride Sizei, pointer *GLvoid)

// var IsAsyncMarkerSGIX func(marker uint) GLboolean
var IsBuffer func(buffer uint) GLboolean
var IsBufferARB func(buffer uint) GLboolean

// var IsBufferResidentNV func(target Enum) GLboolean
var IsEnabled func(cap Enum) GLboolean
var IsEnabledi func(target Enum, index uint) GLboolean
var IsEnabledIndexedEXT func(target Enum, index uint) GLboolean

// var IsFenceAPPLE func(fence uint) GLboolean
// var IsFenceNV func(fence uint) GLboolean
var IsFramebuffer func(framebuffer uint) GLboolean
var IsFramebufferEXT func(framebuffer uint) GLboolean
var IsList func(list uint) GLboolean

// var IsNameAMD func(identifier Enum, name uint) GLboolean
// var IsNamedBufferResidentNV func(buffer uint) GLboolean
// var IsNamedStringARB func(namelen int, name string) GLboolean
// var IsObjectBufferATI func(buffer uint) GLboolean
// var IsOcclusionQueryNV func(id uint) GLboolean
var IsProgram func(program uint) GLboolean
var IsProgramARB func(program uint) GLboolean
var IsProgramNV func(id uint) GLboolean

// var IsProgramPipeline func(pipeline uint) GLboolean
var IsQuery func(id uint) GLboolean
var IsQueryARB func(id uint) GLboolean
var IsRenderbuffer func(renderbuffer uint) GLboolean
var IsRenderbufferEXT func(renderbuffer uint) GLboolean
var IsSampler func(sampler uint) GLboolean
var IsShader func(shader uint) GLboolean
var IsSync func(sync Sync) GLboolean
var IsTexture func(texture uint) GLboolean
var IsTextureEXT func(texture uint) GLboolean
var IsTransformFeedback func(id uint) GLboolean

// var IsTransformFeedbackNV func(id uint) GLboolean
// var IsVariantEnabledEXT func(id uint, cap Enum) GLboolean
var IsVertexArray func(array uint) GLboolean

// var IsVertexArrayAPPLE func(array uint) GLboolean
// var IsVertexAttribEnabledAPPLE func(index uint, pname Enum) GLboolean
// var LightEnviSGIX func(pname Enum, param int)
var Lightf func(light, pname Enum, param float32)
var Lightfv func(light, pname Enum, params *float32)
var Lighti func(light, pname Enum, param int)
var Lightiv func(light, pname Enum, params *int)
var LightModelf func(pname Enum, param float32)
var LightModelfv func(pname Enum, params *float32)
var LightModeli func(pname Enum, param int)
var LightModeliv func(pname Enum, params *int)
var LineStipple func(factor int, pattern uint16)
var LineWidth func(width float32)
var LinkProgram func(program uint)
var LinkProgramARB func(programObj HandleARB)
var ListBase func(base uint)

// var ListParameterfSGIX func(list uint, pname Enum, param float32)
// var ListParameterfvSGIX func(list uint, pname Enum, params *float32)
// var ListParameteriSGIX func(list uint, pname Enum, param int)
// var ListParameterivSGIX func(list uint, pname Enum, params *int)
var LoadIdentity func()

// var LoadIdentityDeformationMapSGIX func(mask GLbitfield)
var LoadMatrixd func(m *float64)
var LoadMatrixf func(m *float32)
var LoadName func(name uint)
var LoadProgramNV func(target Enum, id uint, len Sizei, program *uint8)
var LoadTransposeMatrixd func(m [16]float64)
var LoadTransposeMatrixdARB func(m *float64)
var LoadTransposeMatrixf func(m [16]float32)
var LoadTransposeMatrixfARB func(m *float32)
var LockArraysEXT func(first int, count Sizei)
var LogicOp func(opcode Enum)

// var MakeBufferNonResidentNV func(target Enum)
// var MakeBufferResidentNV func(target, access Enum)
// var MakeNamedBufferNonResidentNV func(buffer uint)
// var MakeNamedBufferResidentNV func(buffer uint, access Enum)
var Map1d func(target Enum, u1, u2 float64, stride, order int, points *float64)
var Map1f func(target Enum, u1, u2 float32, stride, order int, points *float32)
var Map2d func(target Enum, u1, u2 float64, ustride, uorder int, v1, v2 float64, vstride, vorder int, points *float64)
var Map2f func(target Enum, u1, u2 float32, ustride, uorder int, v1, v2 float32, vstride, vorder int, points *float32)
var MapBuffer func(target, access Enum) *GLvoid
var MapBufferARB func(target, access Enum) *GLvoid
var MapBufferRange func(target Enum, offset Intptr, length Sizeiptr, access GLbitfield) *GLvoid

// var MapControlPointsNV func(target Enum, index uint, type_ Enum, ustride, vstride Sizei, uorder, vorder int, packed GLboolean, points *GLvoid)
var MapGrid1d func(un int, u1, u2 float64)
var MapGrid1f func(un int, u1, u2 float32)
var MapGrid2d func(un int, u1, u2 float64, vn int, v1, v2 float64)
var MapGrid2f func(un int, u1, u2 float32, vn int, v1, v2 float32)

// var MapNamedBufferEXT func(buffer uint, access Enum) *GLvoid
// var MapNamedBufferRangeEXT func(buffer uint, offset Intptr, length Sizeiptr, access GLbitfield) *GLvoid
// var MapObjectBufferATI func(buffer uint) *GLvoid
// var MapParameterfvNV func(target, pname Enum, params *float32)
// var MapParameterivNV func(target, pname Enum, params *int)
// var MapVertexAttrib1dAPPLE func(index, size uint, u1, u2 float64, stride, order int, points *float64)
// var MapVertexAttrib1fAPPLE func(index, size uint, u1, u2 float32, stride, order int, points *float32)
// var MapVertexAttrib2dAPPLE func(index, size uint, u1, u2 float64, ustride, uorder int, v1, v2 float64, vstride, vorder int, points *float64)
// var MapVertexAttrib2fAPPLE func(index, size uint, u1, u2 float32, ustride, uorder int, v1, v2 float32, vstride, vorder int, points *float32)
var Materialf func(face, pname Enum, param float32)
var Materialfv func(face, pname Enum, params *float32)
var Materiali func(face, pname Enum, param int)
var Materialiv func(face, pname Enum, params *int)

// var MatrixFrustumEXT func(mode Enum, left, right, bottom, top, zNear, zFar float64)
// var MatrixIndexPointerARB func(size int, type_ Enum, stride Sizei, pointer *GLvoid)
// var MatrixIndexubvARB func(size int, indices *uint8)
// var MatrixIndexuivARB func(size int, indices *uint)
// var MatrixIndexusvARB func(size int, indices *uint16)
// var MatrixLoaddEXT func(mode Enum, m *float64)
// var MatrixLoadfEXT func(mode Enum, m *float32)
// var MatrixLoadIdentityEXT func(mode Enum)
// var MatrixLoadTransposedEXT func(mode Enum, m *float64)
// var MatrixLoadTransposefEXT func(mode Enum, m *float32)
var MatrixMode func(mode Enum)

// var MatrixMultdEXT func(mode Enum, m *float64)
// var MatrixMultfEXT func(mode Enum, m *float32)
// var MatrixMultTransposedEXT func(mode Enum, m *float64)
// var MatrixMultTransposefEXT func(mode Enum, m *float32)
// var MatrixOrthoEXT func(mode Enum, left, right, bottom, top, zNear, zFar float64)
// var MatrixPopEXT func(mode Enum)
// var MatrixPushEXT func(mode Enum)
// var MatrixRotatedEXT func(mode Enum, angle, x, y, z float64)
// var MatrixRotatefEXT func(mode Enum, angle, x, y, z float32)
// var MatrixScaledEXT func(mode Enum, x, y, z float64)
// var MatrixScalefEXT func(mode Enum, x, y, z float32)
// var MatrixTranslatedEXT func(mode Enum, x, y, z float64)
// var MatrixTranslatefEXT func(mode Enum, x, y, z float32)
// var MemoryBarrier func(barriers GLbitfield)
// var MemoryBarrierEXT func(barriers GLbitfield)
var Minmax func(target, internalformat Enum, sink GLboolean)

// var MinmaxEXT func(target, internalformat Enum, sink GLboolean)
// var MinSampleShading func(value clampf)
// var MinSampleShadingARB func(value clampf)
var MultiDrawArrays func(mode Enum, first *int, count *Sizei, primcount Sizei)
var MultiDrawArraysEXT func(mode Enum, first *int, count *Sizei, primcount Sizei)

// var MultiDrawArraysIndirectAMD func(mode Enum, indirect *GLvoid, primcount, stride Sizei)
// var MultiDrawElementArrayAPPLE func(mode Enum, first *int, count *Sizei, primcount Sizei)
var MultiDrawElements func(mode Enum, count *Sizei, type_ Enum, indices **GLvoid, primcount Sizei)
var MultiDrawElementsBaseVertex func(mode Enum, count *Sizei, type_ Enum, indices **GLvoid, primcount Sizei, basevertex *int)
var MultiDrawElementsEXT func(mode Enum, count *Sizei, type_ Enum, indices **GLvoid, primcount Sizei)

// var MultiDrawElementsIndirectAMD func(mode, type_ Enum, indirect *GLvoid, primcount, stride Sizei)
// var MultiDrawRangeElementArrayAPPLE func(mode Enum, start, end uint, first *int, count *Sizei, primcount Sizei)
// var MultiModeDrawArraysIBM func(mode *Enum, first *int, count *Sizei, primcount Sizei, modestride int)
// var MultiModeDrawElementsIBM func(mode *Enum, count *Sizei, type_ Enum, indices **GLvoid, primcount Sizei, modestride int)
// var MultiTexBufferEXT func(texunit, target, internalformat Enum, buffer uint)
var MultiTexCoord1d func(target Enum, s float64)
var MultiTexCoord1dARB func(target Enum, s float64)
var MultiTexCoord1dv func(target Enum, v *float64)
var MultiTexCoord1dvARB func(target Enum, v *float64)
var MultiTexCoord1f func(target Enum, s float32)
var MultiTexCoord1fARB func(target Enum, s float32)
var MultiTexCoord1fv func(target Enum, v *float32)
var MultiTexCoord1fvARB func(target Enum, v *float32)

// var MultiTexCoord1hNV func(target Enum, s GLhalfNV)
// var MultiTexCoord1hvNV func(target Enum, v *GLhalfNV)
var MultiTexCoord1i func(target Enum, s int)
var MultiTexCoord1iARB func(target Enum, s int)
var MultiTexCoord1iv func(target Enum, v *int)
var MultiTexCoord1ivARB func(target Enum, v *int)
var MultiTexCoord1s func(target Enum, s int16)
var MultiTexCoord1sARB func(target Enum, s int16)
var MultiTexCoord1sv func(target Enum, v *int16)
var MultiTexCoord1svARB func(target Enum, v *int16)
var MultiTexCoord2d func(target Enum, s, t float64)
var MultiTexCoord2dARB func(target Enum, s, t float64)
var MultiTexCoord2dv func(target Enum, v *float64)
var MultiTexCoord2dvARB func(target Enum, v *float64)
var MultiTexCoord2f func(target Enum, s, t float32)
var MultiTexCoord2fARB func(target Enum, s, t float32)
var MultiTexCoord2fv func(target Enum, v *float32)
var MultiTexCoord2fvARB func(target Enum, v *float32)

// var MultiTexCoord2hNV func(target Enum, s, t GLhalfNV)
// var MultiTexCoord2hvNV func(target Enum, v *GLhalfNV)
var MultiTexCoord2i func(target Enum, s, t int)
var MultiTexCoord2iARB func(target Enum, s, t int)
var MultiTexCoord2iv func(target Enum, v *int)
var MultiTexCoord2ivARB func(target Enum, v *int)
var MultiTexCoord2s func(target Enum, s, t int16)
var MultiTexCoord2sARB func(target Enum, s, t int16)
var MultiTexCoord2sv func(target Enum, v *int16)
var MultiTexCoord2svARB func(target Enum, v *int16)
var MultiTexCoord3d func(target Enum, s, t, r float64)
var MultiTexCoord3dARB func(target Enum, s, t, r float64)
var MultiTexCoord3dv func(target Enum, v *float64)
var MultiTexCoord3dvARB func(target Enum, v *float64)
var MultiTexCoord3f func(target Enum, s, t, r float32)
var MultiTexCoord3fARB func(target Enum, s, t, r float32)
var MultiTexCoord3fv func(target Enum, v *float32)
var MultiTexCoord3fvARB func(target Enum, v *float32)

// var MultiTexCoord3hNV func(target Enum, s, t, r GLhalfNV)
// var MultiTexCoord3hvNV func(target Enum, v *GLhalfNV)
var MultiTexCoord3i func(target Enum, s, t, r int)
var MultiTexCoord3iARB func(target Enum, s, t, r int)
var MultiTexCoord3iv func(target Enum, v *int)
var MultiTexCoord3ivARB func(target Enum, v *int)
var MultiTexCoord3s func(target Enum, s, t, r int16)
var MultiTexCoord3sARB func(target Enum, s, t, r int16)
var MultiTexCoord3sv func(target Enum, v *int16)
var MultiTexCoord3svARB func(target Enum, v *int16)
var MultiTexCoord4d func(target Enum, s, t, r, q float64)
var MultiTexCoord4dARB func(target Enum, s, t, r, q float64)
var MultiTexCoord4dv func(target Enum, v *float64)
var MultiTexCoord4dvARB func(target Enum, v *float64)
var MultiTexCoord4f func(target Enum, s, t, r, q float32)
var MultiTexCoord4fARB func(target Enum, s, t, r, q float32)
var MultiTexCoord4fv func(target Enum, v *float32)
var MultiTexCoord4fvARB func(target Enum, v *float32)

// var MultiTexCoord4hNV func(target Enum, s, t, r, q GLhalfNV)
// var MultiTexCoord4hvNV func(target Enum, v *GLhalfNV)
var MultiTexCoord4i func(target Enum, s, t, r, q int)
var MultiTexCoord4iARB func(target Enum, s, t, r, q int)
var MultiTexCoord4iv func(target Enum, v *int)
var MultiTexCoord4ivARB func(target Enum, v *int)
var MultiTexCoord4s func(target Enum, s, t, r, q int16)
var MultiTexCoord4sARB func(target Enum, s, t, r, q int16)
var MultiTexCoord4sv func(target Enum, v *int16)
var MultiTexCoord4svARB func(target Enum, v *int16)
var MultiTexCoordP1ui func(texture, type_ Enum, coords uint)
var MultiTexCoordP1uiv func(texture, type_ Enum, coords *uint)
var MultiTexCoordP2ui func(texture, type_ Enum, coords uint)
var MultiTexCoordP2uiv func(texture, type_ Enum, coords *uint)
var MultiTexCoordP3ui func(texture, type_ Enum, coords uint)
var MultiTexCoordP3uiv func(texture, type_ Enum, coords *uint)
var MultiTexCoordP4ui func(texture, type_ Enum, coords uint)
var MultiTexCoordP4uiv func(texture, type_ Enum, coords *uint)

// var MultiTexCoordPointerEXT func(texunit Enum, size int, type_ Enum, stride Sizei, pointer *GLvoid)
// var MultiTexEnvfEXT func(texunit, target, pname Enum, param float32)
// var MultiTexEnvfvEXT func(texunit, target, pname Enum, params *float32)
// var MultiTexEnviEXT func(texunit, target, pname Enum, param int)
// var MultiTexEnvivEXT func(texunit, target, pname Enum, params *int)
// var MultiTexGendEXT func(texunit, coord, pname Enum, param float64)
// var MultiTexGendvEXT func(texunit, coord, pname Enum, params *float64)
// var MultiTexGenfEXT func(texunit, coord, pname Enum, param float32)
// var MultiTexGenfvEXT func(texunit, coord, pname Enum, params *float32)
// var MultiTexGeniEXT func(texunit, coord, pname Enum, param int)
// var MultiTexGenivEXT func(texunit, coord, pname Enum, params *int)
// var MultiTexImage1DEXT func(texunit, target Enum, level int, internalformat Enum, width Sizei, border int, format, type_ Enum, pixels *GLvoid)
// var MultiTexImage2DEXT func(texunit, target Enum, level int, internalformat Enum, width, height Sizei, border int, format, type_ Enum, pixels *GLvoid)
// var MultiTexImage3DEXT func(texunit, target Enum, level int, internalformat Enum, width, height, depth Sizei, border int, format, type_ Enum, pixels *GLvoid)
// var MultiTexParameterfEXT func(texunit, target, pname Enum, param float32)
// var MultiTexParameterfvEXT func(texunit, target, pname Enum, params *float32)
// var MultiTexParameteriEXT func(texunit, target, pname Enum, param int)
// var MultiTexParameterIivEXT func(texunit, target, pname Enum, params *int)
// var MultiTexParameterIuivEXT func(texunit, target, pname Enum, params *uint)
// var MultiTexParameterivEXT func(texunit, target, pname Enum, params *int)
// var MultiTexRenderbufferEXT func(texunit, target Enum, renderbuffer uint)
// var MultiTexSubImage1DEXT func(texunit, target Enum, level, xoffset int, width Sizei, format, type_ Enum, pixels *GLvoid)
// var MultiTexSubImage2DEXT func(texunit, target Enum, level, xoffset, yoffset int, width, height Sizei, format, type_ Enum, pixels *GLvoid)
// var MultiTexSubImage3DEXT func(texunit, target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format, type_ Enum, pixels *GLvoid)
var MultMatrixd func(m *float64)
var MultMatrixf func(m *float32)
var MultTransposeMatrixd func(m [16]float64)
var MultTransposeMatrixdARB func(m *float64)
var MultTransposeMatrixf func(m [16]float32)
var MultTransposeMatrixfARB func(m *float32)

// var NamedBufferDataEXT func(buffer uint, size Sizeiptr, data *GLvoid, usage Enum)
// var NamedBufferSubDataEXT func(buffer uint, offset Intptr, size Sizeiptr, data *GLvoid)
// var NamedCopyBufferSubDataEXT func(readBuffer, writeBuffer uint, readOffset, writeOffset Intptr, size Sizeiptr)
// var NamedFramebufferRenderbufferEXT func(framebuffer uint, attachment, renderbuffertarget Enum, renderbuffer uint)
// var NamedFramebufferTexture1DEXT func(framebuffer uint, attachment, textarget Enum, texture uint, level int)
// var NamedFramebufferTexture2DEXT func(framebuffer uint, attachment, textarget Enum, texture uint, level int)
// var NamedFramebufferTexture3DEXT func(framebuffer uint, attachment, textarget Enum, texture uint, level, zoffset int)
// var NamedFramebufferTextureEXT func(framebuffer uint, attachment Enum, texture uint, level int)
// var NamedFramebufferTextureFaceEXT func(framebuffer uint, attachment Enum, texture uint, level int, face Enum)
// var NamedFramebufferTextureLayerEXT func(framebuffer uint, attachment Enum, texture uint, level, layer int)
// var NamedProgramLocalParameter4dEXT func(program uint, target Enum, index uint, x, y, z, w float64)
// var NamedProgramLocalParameter4dvEXT func(program uint, target Enum, index uint, params *float64)
// var NamedProgramLocalParameter4fEXT func(program uint, target Enum, index uint, x, y, z, w float32)
// var NamedProgramLocalParameter4fvEXT func(program uint, target Enum, index uint, params *float32)
// var NamedProgramLocalParameterI4iEXT func(program uint, target Enum, index uint, x, y, z, w int)
// var NamedProgramLocalParameterI4ivEXT func(program uint, target Enum, index uint, params *int)
// var NamedProgramLocalParameterI4uiEXT func(program uint, target Enum, index, x, y, z, w uint)
// var NamedProgramLocalParameterI4uivEXT func(program uint, target Enum, index uint, params *uint)
// var NamedProgramLocalParameters4fvEXT func(program uint, target Enum, index uint, count Sizei, params *float32)
// var NamedProgramLocalParametersI4ivEXT func(program uint, target Enum, index uint, count Sizei, params *int)
// var NamedProgramLocalParametersI4uivEXT func(program uint, target Enum, index uint, count Sizei, params *uint)
// var NamedProgramStringEXT func(program uint, target, format Enum, len Sizei, string *GLvoid)
// var NamedRenderbufferStorageEXT func(renderbuffer uint, internalformat Enum, width, height Sizei)
// var NamedRenderbufferStorageMultisampleCoverageEXT func(renderbuffer uint, coverageSamples, colorSamples Sizei, internalformat Enum, width, height Sizei)
// var NamedRenderbufferStorageMultisampleEXT func(renderbuffer uint, samples Sizei, internalformat Enum, width, height Sizei)
// var NamedStringARB func(type_ Enum, namelen int, name string, stringlen int, string string)
var NewList func(list uint, mode Enum)

// var NewObjectBufferATI func(size Sizei, pointer *GLvoid, usage Enum) uint
var Normal3b func(hnx, ny, nz GLbyte)
var Normal3bv func(v *GLbyte)
var Normal3d func(hnx, ny, nz float64)
var Normal3dv func(v *float64)
var Normal3f func(hnx, ny, nz float32)
var Normal3fv func(v *float32)

// var Normal3fVertex3fSUN func(nx, ny, nz, x, y, z float32)
// var Normal3fVertex3fvSUN func(n, v *float32)
// var Normal3hNV func(nx, ny, nz GLhalfNV)
// var Normal3hvNV func(v *GLhalfNV)
var Normal3i func(hnx, ny, nz int)
var Normal3iv func(v *int)
var Normal3s func(hnx, ny, nz int16)
var Normal3sv func(v *int16)

// var NormalFormatNV func(type_ Enum, stride Sizei)
var NormalP3ui func(type_ Enum, coords uint)
var NormalP3uiv func(type_ Enum, coords *uint)
var NormalPointer func(type_ Enum, stride Sizei, ptr *GLvoid)
var NormalPointerEXT func(type_ Enum, stride, count Sizei, pointer *GLvoid)

// var NormalPointerListIBM func(type_ Enum, stride int, pointer **GLvoid, ptrstride int)
// var NormalPointervINTEL func(type_ Enum, pointer **GLvoid)
// var NormalStream3bATI func(stream Enum, nx, ny, nz GLbyte)
// var NormalStream3bvATI func(stream Enum, coords *GLbyte)
// var NormalStream3dATI func(stream Enum, nx, ny, nz float64)
// var NormalStream3dvATI func(stream Enum, coords *float64)
// var NormalStream3fATI func(stream Enum, nx, ny, nz float32)
// var NormalStream3fvATI func(stream Enum, coords *float32)
// var NormalStream3iATI func(stream Enum, nx, ny, nz int)
// var NormalStream3ivATI func(stream Enum, coords *int)
// var NormalStream3sATI func(stream Enum, nx, ny, nz int16)
// var NormalStream3svATI func(stream Enum, coords *int16)
var ObjectPurgeableAPPLE func(objectType Enum, name uint, option Enum) Enum
var ObjectUnpurgeableAPPLE func(objectType Enum, name uint, option Enum) Enum
var Ortho func(hhleft, right, bottom, top, near_val, far_val float64)
var PassTexCoordATI func(dst, coord uint, swizzle Enum)
var PassThrough func(token float32)

// var PatchParameterfv func(pname Enum, values *float32)
// var PatchParameteri func(pname Enum, value int)
var PauseTransformFeedback func()

// var PauseTransformFeedbackNV func()
// var PixelDataRangeNV func(target Enum, length Sizei, pointer *GLvoid)
var PixelMapfv func(map_ Enum, mapsize Sizei, values *float32)
var PixelMapuiv func(map_ Enum, mapsize Sizei, values *uint)
var PixelMapusv func(map_ Enum, mapsize Sizei, values *uint16)
var PixelStoref func(pname Enum, param float32)
var PixelStorei func(pname Enum, param int)

// var PixelTexGenParameterfSGIS func(pname Enum, param float32)
// var PixelTexGenParameterfvSGIS func(pname Enum, params *float32)
// var PixelTexGenParameteriSGIS func(pname Enum, param int)
// var PixelTexGenParameterivSGIS func(pname Enum, params *int)
// var PixelTexGenSGIX func(mode Enum)
var PixelTransferf func(pname Enum, param float32)
var PixelTransferi func(pname Enum, param int)

// var PixelTransformParameterfEXT func(target, pname Enum, param float32)
// var PixelTransformParameterfvEXT func(target, pname Enum, params *float32)
// var PixelTransformParameteriEXT func(target, pname Enum, param int)
// var PixelTransformParameterivEXT func(target, pname Enum, params *int)
var PixelZoom func(xfactor, yfactor float32)

// var PNTrianglesfATI func(pname Enum, param float32)
// var PNTrianglesiATI func(pname Enum, param int)
var PointParameterf func(pname Enum, param float32)
var PointParameterfARB func(pname Enum, param float32)
var PointParameterfEXT func(pname Enum, param float32)

// var PointParameterfSGIS func(pname Enum, param float32)
var PointParameterfv func(pname Enum, params *float32)
var PointParameterfvARB func(pname Enum, params *float32)
var PointParameterfvEXT func(pname Enum, params *float32)

// var PointParameterfvSGIS func(pname Enum, params *float32)
var PointParameteri func(pname Enum, param int)
var PointParameteriNV func(pname Enum, param int)
var PointParameteriv func(pname Enum, params *int)
var PointParameterivNV func(pname Enum, params *int)
var PointSize func(size float32)

// var PollAsyncSGIX func(markerp *uint) int
// var PollInstrumentsSGIX func(marker_p *int) int
var PolygonMode func(face, mode Enum)
var PolygonOffset func(factor, units float32)
var PolygonOffsetEXT func(factor, bias float32)
var PolygonStipple func(mask *uint8)
var PopAttrib func()
var PopClientAttrib func()
var PopMatrix func()
var PopName func()

// var PresentFrameDualFillNV func(video_slot, minPresentTime uint64EXT, beginPresentTimeId, presentDurationId uint, type_, target0 Enum, fill0 uint, target1 Enum, fill1 uint, target2 Enum, fill2 uint, target3 Enum, fill3 uint)
// var PresentFrameKeyedNV func(video_slot, minPresentTime uint64EXT, beginPresentTimeId, presentDurationId uint, type_, target0 Enum, fill0, key0 uint, target1 Enum, fill1, key1 uint)
var PrimitiveRestartIndex func(index uint)
var PrimitiveRestartIndexNV func(index uint)
var PrimitiveRestartNV func()
var PrioritizeTextures func(n Sizei, textures *uint, priorities *Clampf)
var PrioritizeTexturesEXT func(n Sizei, textures *uint, priorities *Clampf)

// var ProgramBinary func(program uint, binaryFormat Enum, binary *GLvoid, length Sizei)
// var ProgramBufferParametersfvNV func(target Enum, buffer, index uint, count Sizei, params *float32)
// var ProgramBufferParametersIivNV func(target Enum, buffer, index uint, count Sizei, params *int)
// var ProgramBufferParametersIuivNV func(target Enum, buffer, index uint, count Sizei, params *uint)
// var ProgramCallbackMESA func(target Enum, callback GLprogramcallbackMESA, data *GLvoid)
var ProgramEnvParameter4dARB func(target Enum, index uint, x, y, z, w float64)
var ProgramEnvParameter4dvARB func(target Enum, index uint, params *float64)
var ProgramEnvParameter4fARB func(target Enum, index uint, x, y, z, w float32)
var ProgramEnvParameter4fvARB func(target Enum, index uint, params *float32)

// var ProgramEnvParameterI4iNV func(target Enum, index uint, x, y, z, w int)
// var ProgramEnvParameterI4ivNV func(target Enum, index uint, params *int)
// var ProgramEnvParameterI4uiNV func(target Enum, index, x, y, z, w uint)
// var ProgramEnvParameterI4uivNV func(target Enum, index uint, params *uint)
// var ProgramEnvParameters4fvEXT func(target Enum, index uint, count Sizei, params *float32)
// var ProgramEnvParametersI4ivNV func(target Enum, index uint, count Sizei, params *int)
// var ProgramEnvParametersI4uivNV func(target Enum, index uint, count Sizei, params *uint)
var ProgramLocalParameter4dARB func(target Enum, index uint, x, y, z, w float64)
var ProgramLocalParameter4dvARB func(target Enum, index uint, params *float64)
var ProgramLocalParameter4fARB func(target Enum, index uint, x, y, z, w float32)
var ProgramLocalParameter4fvARB func(target Enum, index uint, params *float32)

// var ProgramLocalParameterI4iNV func(target Enum, index uint, x, y, z, w int)
// var ProgramLocalParameterI4ivNV func(target Enum, index uint, params *int)
// var ProgramLocalParameterI4uiNV func(target Enum, index, x, y, z, w uint)
// var ProgramLocalParameterI4uivNV func(target Enum, index uint, params *uint)
// var ProgramLocalParameters4fvEXT func(target Enum, index uint, count Sizei, params *float32)
// var ProgramLocalParametersI4ivNV func(target Enum, index uint, count Sizei, params *int)
// var ProgramLocalParametersI4uivNV func(target Enum, index uint, count Sizei, params *uint)
var ProgramNamedParameter4dNV func(id uint, len Sizei, name *uint8, x, y, z, w float64)
var ProgramNamedParameter4dvNV func(id uint, len Sizei, name *uint8, v *float64)
var ProgramNamedParameter4fNV func(id uint, len Sizei, name *uint8, x, y, z, w float32)
var ProgramNamedParameter4fvNV func(id uint, len Sizei, name *uint8, v *float32)
var ProgramParameter4dNV func(target Enum, index uint, x, y, z, w float64)
var ProgramParameter4dvNV func(target Enum, index uint, v *float64)
var ProgramParameter4fNV func(target Enum, index uint, x, y, z, w float32)
var ProgramParameter4fvNV func(target Enum, index uint, v *float32)

// var ProgramParameteri func(program uint, pname Enum, value int)
var ProgramParameteriARB func(program uint, pname Enum, value int)

// var ProgramParameteriEXT func(program uint, pname Enum, value int)
var ProgramParameters4dvNV func(target Enum, index uint, count Sizei, v *float64)
var ProgramParameters4fvNV func(target Enum, index uint, count Sizei, v *float32)
var ProgramStringARB func(target, format Enum, len Sizei, string *GLvoid)

// var ProgramSubroutineParametersuivNV func(target Enum, count Sizei, params *uint)
// var ProgramUniform1d func(program uint, location int, v0 float64)
// var ProgramUniform1dEXT func(program uint, location int, x float64)
// var ProgramUniform1dv func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform1dvEXT func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform1f func(program uint, location int, v0 float32)
// var ProgramUniform1fEXT func(program uint, location int, v0 float32)
// var ProgramUniform1fv func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform1fvEXT func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform1i func(program uint, location, v0 int)
// var ProgramUniform1i64NV func(program uint, location, x int64EXT)
// var ProgramUniform1i64vNV func(program uint, location int, count Sizei, value *int64EXT)
// var ProgramUniform1iEXT func(program uint, location, v0 int)
// var ProgramUniform1iv func(program uint, location int, count Sizei, value *int)
// var ProgramUniform1ivEXT func(program uint, location int, count Sizei, value *int)
// var ProgramUniform1ui func(program uint, location int, v0 uint)
// var ProgramUniform1ui64NV func(program uint, location int, x uint64EXT)
// var ProgramUniform1ui64vNV func(program uint, location int, count Sizei, value *uint64EXT)
// var ProgramUniform1uiEXT func(program uint, location int, v0 uint)
// var ProgramUniform1uiv func(program uint, location int, count Sizei, value *uint)
// var ProgramUniform1uivEXT func(program uint, location int, count Sizei, value *uint)
// var ProgramUniform2d func(program uint, location int, v0, v1 float64)
// var ProgramUniform2dEXT func(program uint, location int, x, y float64)
// var ProgramUniform2dv func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform2dvEXT func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform2f func(program uint, location int, v0, v1 float32)
// var ProgramUniform2fEXT func(program uint, location int, v0, v1 float32)
// var ProgramUniform2fv func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform2fvEXT func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform2i func(program uint, location, v0, v1 int)
// var ProgramUniform2i64NV func(program uint, location, x, y int64EXT)
// var ProgramUniform2i64vNV func(program uint, location int, count Sizei, value *int64EXT)
// var ProgramUniform2iEXT func(program uint, location, v0, v1 int)
// var ProgramUniform2iv func(program uint, location int, count Sizei, value *int)
// var ProgramUniform2ivEXT func(program uint, location int, count Sizei, value *int)
// var ProgramUniform2ui func(program uint, location int, v0, v1 uint)
// var ProgramUniform2ui64NV func(program uint, location int, x, y uint64EXT)
// var ProgramUniform2ui64vNV func(program uint, location int, count Sizei, value *uint64EXT)
// var ProgramUniform2uiEXT func(program uint, location int, v0, v1 uint)
// var ProgramUniform2uiv func(program uint, location int, count Sizei, value *uint)
// var ProgramUniform2uivEXT func(program uint, location int, count Sizei, value *uint)
// var ProgramUniform3d func(program uint, location int, v0, v1, v2 float64)
// var ProgramUniform3dEXT func(program uint, location int, x, y, z float64)
// var ProgramUniform3dv func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform3dvEXT func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform3f func(program uint, location int, v0, v1, v2 float32)
// var ProgramUniform3fEXT func(program uint, location int, v0, v1, v2 float32)
// var ProgramUniform3fv func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform3fvEXT func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform3i func(program uint, location, v0, v1, v2 int)
// var ProgramUniform3i64NV func(program uint, location, x, y, z int64EXT)
// var ProgramUniform3i64vNV func(program uint, location int, count Sizei, value *int64EXT)
// var ProgramUniform3iEXT func(program uint, location, v0, v1, v2 int)
// var ProgramUniform3iv func(program uint, location int, count Sizei, value *int)
// var ProgramUniform3ivEXT func(program uint, location int, count Sizei, value *int)
// var ProgramUniform3ui func(program uint, location int, v0, v1, v2 uint)
// var ProgramUniform3ui64NV func(program uint, location int, x, y, z uint64EXT)
// var ProgramUniform3ui64vNV func(program uint, location int, count Sizei, value *uint64EXT)
// var ProgramUniform3uiEXT func(program uint, location int, v0, v1, v2 uint)
// var ProgramUniform3uiv func(program uint, location int, count Sizei, value *uint)
// var ProgramUniform3uivEXT func(program uint, location int, count Sizei, value *uint)
// var ProgramUniform4d func(program uint, location int, v0, v1, v2, v3 float64)
// var ProgramUniform4dEXT func(program uint, location int, x, y, z, w float64)
// var ProgramUniform4dv func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform4dvEXT func(program uint, location int, count Sizei, value *float64)
// var ProgramUniform4f func(program uint, location int, v0, v1, v2, v3 float32)
// var ProgramUniform4fEXT func(program uint, location int, v0, v1, v2, v3 float32)
// var ProgramUniform4fv func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform4fvEXT func(program uint, location int, count Sizei, value *float32)
// var ProgramUniform4i func(program uint, location, v0, v1, v2, v3 int)
// var ProgramUniform4i64NV func(program uint, location, x, y, z, w int64EXT)
// var ProgramUniform4i64vNV func(program uint, location int, count Sizei, value *int64EXT)
// var ProgramUniform4iEXT func(program uint, location, v0, v1, v2, v3 int)
// var ProgramUniform4iv func(program uint, location int, count Sizei, value *int)
// var ProgramUniform4ivEXT func(program uint, location int, count Sizei, value *int)
// var ProgramUniform4ui func(program uint, location int, v0, v1, v2, v3 uint)
// var ProgramUniform4ui64NV func(program uint, location int, x, y, z, w uint64EXT)
// var ProgramUniform4ui64vNV func(program uint, location int, count Sizei, value *uint64EXT)
// var ProgramUniform4uiEXT func(program uint, location int, v0, v1, v2, v3 uint)
// var ProgramUniform4uiv func(program uint, location int, count Sizei, value *uint)
// var ProgramUniform4uivEXT func(program uint, location int, count Sizei, value *uint)
// var ProgramUniformMatrix2dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix2dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix2fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix2fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix2x3dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix2x3dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix2x3fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix2x3fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix2x4dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix2x4dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix2x4fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix2x4fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix3dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix3dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix3fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix3fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix3x2dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix3x2dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix3x2fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix3x2fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix3x4dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix3x4dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix3x4fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix3x4fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix4dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix4dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix4fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix4fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix4x2dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix4x2dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix4x2fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix4x2fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix4x3dv func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix4x3dvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float64)
// var ProgramUniformMatrix4x3fv func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformMatrix4x3fvEXT func(program uint, location int, count Sizei, transpose GLboolean, value *float32)
// var ProgramUniformui64NV func(program uint, location int, value uint64EXT)
// var ProgramUniformui64vNV func(program uint, location int, count Sizei, value *uint64EXT)
// var ProgramVertexLimitNV func(target Enum, limit int)
var ProvokingVertex func(mode Enum)
var ProvokingVertexEXT func(mode Enum)
var PushAttrib func(mask GLbitfield)
var PushClientAttrib func(mask GLbitfield)

// var PushClientAttribDefaultEXT func(mask GLbitfield)
var PushMatrix func()
var PushName func(name uint)

// var QueryCounter func(id uint, target Enum)
var RasterPos2d func(x, y float64)
var RasterPos2dv func(v *float64)
var RasterPos2f func(x, y float32)
var RasterPos2fv func(v *float32)
var RasterPos2i func(x, y int)
var RasterPos2iv func(v *int)
var RasterPos2s func(x, y int16)
var RasterPos2sv func(v *int16)
var RasterPos3d func(hx, y, z float64)
var RasterPos3dv func(v *float64)
var RasterPos3f func(hx, y, z float32)
var RasterPos3fv func(v *float32)
var RasterPos3i func(hx, y, z int)
var RasterPos3iv func(v *int)
var RasterPos3s func(hx, y, z int16)
var RasterPos3sv func(v *int16)
var RasterPos4d func(hx, y, z, w float64)
var RasterPos4dv func(v *float64)
var RasterPos4f func(hx, y, z, w float32)
var RasterPos4fv func(v *float32)
var RasterPos4i func(hx, y, z, w int)
var RasterPos4iv func(v *int)
var RasterPos4s func(hx, y, z, w int16)
var RasterPos4sv func(v *int16)
var ReadBuffer func(mode Enum)

// var ReadInstrumentsSGIX func(marker int)
var ReadnPixelsARB func(x, y int, width, height Sizei, format, type_ Enum, bufSize Sizei, data *GLvoid)
var ReadPixels func(x, y int, width, height Sizei, format, type_ Enum, pixels *GLvoid)
var Rectd func(hx1, y1, x2, y2 float64)
var Rectdv func(v1, v2 *float64)
var Rectf func(hx1, y1, x2, y2 float32)
var Rectfv func(v1, v2 *float32)
var Recti func(hx1, y1, x2, y2 int)
var Rectiv func(v1, v2 *int)
var Rects func(hx1, y1, x2, y2 int16)
var Rectsv func(v1, v2 *int16)

// var ReferencePlaneSGIX func(equation *float64)
var ReleaseShaderCompiler func()
var RenderbufferStorage func(target, internalformat Enum, width, height Sizei)
var RenderbufferStorageEXT func(target, internalformat Enum, width, height Sizei)
var RenderbufferStorageMultisample func(target Enum, samples Sizei, internalformat Enum, width, height Sizei)

// var RenderbufferStorageMultisampleCoverageNV func(target Enum, coverageSamples, colorSamples Sizei, internalformat Enum, width, height Sizei)
var RenderbufferStorageMultisampleEXT func(target Enum, samples Sizei, internalformat Enum, width, height Sizei)
var RenderMode func(mode Enum) int

// var ReplacementCodePointerSUN func(type_ Enum, stride Sizei, pointer **GLvoid)
// var ReplacementCodeubSUN func(code uint8)
// var ReplacementCodeubvSUN func(code *uint8)
// var ReplacementCodeuiColor3fVertex3fSUN func(rc uint, r, g, b, x, y, z float32)
// var ReplacementCodeuiColor3fVertex3fvSUN func(rc *uint, c, v *float32)
// var ReplacementCodeuiColor4fNormal3fVertex3fSUN func(rc uint, r, g, b, a, nx, ny, nz, x, y, z float32)
// var ReplacementCodeuiColor4fNormal3fVertex3fvSUN func(rc *uint, c, n, v *float32)
// var ReplacementCodeuiColor4ubVertex3fSUN func(rc uint, r, g, b, a uint8, x, y, z float32)
// var ReplacementCodeuiColor4ubVertex3fvSUN func(rc *uint, c *uint8, v *float32)
// var ReplacementCodeuiNormal3fVertex3fSUN func(rc uint, nx, ny, nz, x, y, z float32)
// var ReplacementCodeuiNormal3fVertex3fvSUN func(rc *uint, n, v *float32)
// var ReplacementCodeuiSUN func(code uint)
// var ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN func(rc uint, s, t, r, g, b, a, nx, ny, nz, x, y, z float32)
// var ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN func(rc *uint, tc, c, n, v *float32)
// var ReplacementCodeuiTexCoord2fNormal3fVertex3fSUN func(rc uint, s, t, nx, ny, nz, x, y, z float32)
// var ReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN func(rc *uint, tc, n, v *float32)
// var ReplacementCodeuiTexCoord2fVertex3fSUN func(rc uint, s, t, x, y, z float32)
// var ReplacementCodeuiTexCoord2fVertex3fvSUN func(rc *uint, tc, v *float32)
// var ReplacementCodeuiVertex3fSUN func(rc uint, x, y, z float32)
// var ReplacementCodeuiVertex3fvSUN func(rc *uint, v *float32)
// var ReplacementCodeuivSUN func(code *uint)
// var ReplacementCodeusSUN func(code uint16)
// var ReplacementCodeusvSUN func(code *uint16)
var RequestResidentProgramsNV func(n Sizei, programs *uint)
var ResetHistogram func(target Enum)

// var ResetHistogramEXT func(target Enum)
var ResetMinmax func(target Enum)

// var ResetMinmaxEXT func(target Enum)
var ResizeBuffersMESA func()
var ResumeTransformFeedback func()

// var ResumeTransformFeedbackNV func()
var Rotated func(hangle, x, y, z float64)
var Rotatef func(hangle, x, y, z float32)
var SampleCoverage func(value Clampf, invert GLboolean)
var SampleCoverageARB func(value Clampf, invert GLboolean)
var SampleMapATI func(dst, interp uint, swizzle Enum)

// var SampleMaskEXT func(value clampf, invert GLboolean)
// var SampleMaski func(index uint, mask GLbitfield)
// var SampleMaskIndexedNV func(index uint, mask GLbitfield)
// var SampleMaskSGIS func(value clampf, invert GLboolean)
// var SamplePatternEXT func(pattern Enum)
// var SamplePatternSGIS func(pattern Enum)
var SamplerParameterf func(sampler uint, pname Enum, param float32)
var SamplerParameterfv func(sampler uint, pname Enum, param *float32)
var SamplerParameteri func(sampler uint, pname Enum, param int)
var SamplerParameterIiv func(sampler uint, pname Enum, param *int)
var SamplerParameterIuiv func(sampler uint, pname Enum, param *uint)
var SamplerParameteriv func(sampler uint, pname Enum, param *int)
var Scaled func(hx, y, z float64)
var Scalef func(hx, y, z float32)
var Scissor func(x, y int, width, height Sizei)

// var ScissorArrayv func(first uint, count Sizei, v *int)
// var ScissorIndexed func(index uint, left, bottom int, width, height Sizei)
// var ScissorIndexedv func(index uint, v *int)
var SecondaryColor3b func(red, green, blue GLbyte)
var SecondaryColor3bEXT func(red, green, blue GLbyte)
var SecondaryColor3bv func(v *GLbyte)
var SecondaryColor3bvEXT func(v *GLbyte)
var SecondaryColor3d func(red, green, blue float64)
var SecondaryColor3dEXT func(red, green, blue float64)
var SecondaryColor3dv func(v *float64)
var SecondaryColor3dvEXT func(v *float64)
var SecondaryColor3f func(red, green, blue float32)
var SecondaryColor3fEXT func(red, green, blue float32)
var SecondaryColor3fv func(v *float32)
var SecondaryColor3fvEXT func(v *float32)

// var SecondaryColor3hNV func(red, green, blue GLhalfNV)
// var SecondaryColor3hvNV func(v *GLhalfNV)
var SecondaryColor3i func(red, green, blue int)
var SecondaryColor3iEXT func(red, green, blue int)
var SecondaryColor3iv func(v *int)
var SecondaryColor3ivEXT func(v *int)
var SecondaryColor3s func(red, green, blue int16)
var SecondaryColor3sEXT func(red, green, blue int16)
var SecondaryColor3sv func(v *int16)
var SecondaryColor3svEXT func(v *int16)
var SecondaryColor3ub func(red, green, blue uint8)
var SecondaryColor3ubEXT func(red, green, blue uint8)
var SecondaryColor3ubv func(v *uint8)
var SecondaryColor3ubvEXT func(v *uint8)
var SecondaryColor3ui func(red, green, blue uint)
var SecondaryColor3uiEXT func(red, green, blue uint)
var SecondaryColor3uiv func(v *uint)
var SecondaryColor3uivEXT func(v *uint)
var SecondaryColor3us func(red, green, blue uint16)
var SecondaryColor3usEXT func(red, green, blue uint16)
var SecondaryColor3usv func(v *uint16)
var SecondaryColor3usvEXT func(v *uint16)

// var SecondaryColorFormatNV func(size int, type_ Enum, stride Sizei)
var SecondaryColorP3ui func(type_ Enum, color uint)
var SecondaryColorP3uiv func(type_ Enum, color *uint)
var SecondaryColorPointer func(size int, type_ Enum, stride Sizei, pointer *GLvoid)
var SecondaryColorPointerEXT func(size int, type_ Enum, stride Sizei, pointer *GLvoid)

// var SecondaryColorPointerListIBM func(size int, type_ Enum, stride int, pointer **GLvoid, ptrstride int)
var SelectBuffer func(size Sizei, buffer *uint)

// var SelectPerfMonitorCountersAMD func(monitor uint, enable GLboolean, group uint, numCounters int, counterList *uint)
var SeparableFilter2D func(target, internalformat Enum, width, height Sizei, format, type_ Enum, row, column *GLvoid)

// var SeparableFilter2DEXT func(target, internalformat Enum, width, height Sizei, format, type_ Enum, row, column *GLvoid)
// var SetFenceAPPLE func(fence uint)
// var SetFenceNV func(fence uint, condition Enum)
var SetFragmentShaderConstantATI func(dst uint, value *float32)

// var SetInvariantEXT func(id uint, type_ Enum, addr *GLvoid)
// var SetLocalConstantEXT func(id uint, type_ Enum, addr *GLvoid)
// var SetMultisamplefvAMD func(pname Enum, index uint, val *float32)
var ShadeModel func(mode Enum)
var ShaderBinary func(count Sizei, shaders *uint, binaryformat Enum, binary *GLvoid, length Sizei)

// var ShaderOp1EXT func(op Enum, res, arg1 uint)
// var ShaderOp2EXT func(op Enum, res, arg1, arg2 uint)
// var ShaderOp3EXT func(op Enum, res, arg1, arg2, arg3 uint)
var ShaderSource func(shader uint, count Sizei, GLchar **string, length *int)
var ShaderSourceARB func(shaderObj HandleARB, count Sizei, GLcharARB **string, length *int)

// var SharpenTexFuncSGIS func(target Enum, n Sizei, points *float32)
// var SpriteParameterfSGIX func(pname Enum, param float32)
// var SpriteParameterfvSGIX func(pname Enum, params *float32)
// var SpriteParameteriSGIX func(pname Enum, param int)
// var SpriteParameterivSGIX func(pname Enum, params *int)
// var StartInstrumentsSGIX func()
// var StencilClearTagEXT func(stencilTagBits Sizei, stencilClearTag uint)
var StencilFunc func(func_ Enum, ref int, mask uint)
var StencilFuncSeparate func(face, func_ Enum, ref int, mask uint)

// var StencilFuncSeparateATI func(frontfunc_, backfunc_ Enum, ref int, mask uint)
var StencilMask func(mask uint)
var StencilMaskSeparate func(face Enum, mask uint)
var StencilOp func(hfail, zfail, zpass Enum)
var StencilOpSeparate func(face, sfail, dpfail, dppass Enum)

// var StencilOpSeparateATI func(face, sfail, dpfail, dppass Enum)
// var StopInstrumentsSGIX func(marker int)
// var StringMarkerGREMEDY func(len Sizei, string *GLvoid)
// var SwizzleEXT func(res, in uint, outX, outY, outZ, outW Enum)
// var TagSampleBufferSGIX func()
// var Tangent3bEXT func(tx, ty, tz GLbyte)
// var Tangent3bvEXT func(v *GLbyte)
// var Tangent3dEXT func(tx, ty, tz float64)
// var Tangent3dvEXT func(v *float64)
// var Tangent3fEXT func(tx, ty, tz float32)
// var Tangent3fvEXT func(v *float32)
// var Tangent3iEXT func(tx, ty, tz int)
// var Tangent3ivEXT func(v *int)
// var Tangent3sEXT func(tx, ty, tz int16)
// var Tangent3svEXT func(v *int16)
// var TangentPointerEXT func(type_ Enum, stride Sizei, pointer *GLvoid)
// var TbufferMask3DFX func(mask uint)
// var TessellationFactorAMD func(factor float32)
// var TessellationModeAMD func(mode Enum)
// var TestFenceAPPLE func(fence uint) GLboolean
// var TestFenceNV func(fence uint) GLboolean
// var TestObjectAPPLE func(object Enum, name uint) GLboolean
var TexBuffer func(target, internalformat Enum, buffer uint)
var TexBufferARB func(target, internalformat Enum, buffer uint)

// var TexBufferEXT func(target, internalformat Enum, buffer uint)
var TexBumpParameterfvATI func(pname Enum, param *float32)
var TexBumpParameterivATI func(pname Enum, param *int)
var TexCoord1d func(s float64)
var TexCoord1dv func(v *float64)
var TexCoord1f func(s float32)
var TexCoord1fv func(v *float32)

// var TexCoord1hNV func(s GLhalfNV)
// var TexCoord1hvNV func(v *GLhalfNV)
var TexCoord1i func(s int)
var TexCoord1iv func(v *int)
var TexCoord1s func(s int16)
var TexCoord1sv func(v *int16)
var TexCoord2d func(s, t float64)
var TexCoord2dv func(v *float64)
var TexCoord2f func(s, t float32)

// var TexCoord2fColor3fVertex3fSUN func(s, t, r, g, b, x, y, z float32)
// var TexCoord2fColor3fVertex3fvSUN func(tc, c, v *float32)
// var TexCoord2fColor4fNormal3fVertex3fSUN func(s, t, r, g, b, a, nx, ny, nz, x, y, z float32)
// var TexCoord2fColor4fNormal3fVertex3fvSUN func(tc, c, n, v *float32)
// var TexCoord2fColor4ubVertex3fSUN func(s, t float32, r, g, b, a uint8, x, y, z float32)
// var TexCoord2fColor4ubVertex3fvSUN func(tc *float32, c *uint8, v *float32)
// var TexCoord2fNormal3fVertex3fSUN func(s, t, nx, ny, nz, x, y, z float32)
// var TexCoord2fNormal3fVertex3fvSUN func(tc, n, v *float32)
var TexCoord2fv func(v *float32)

// var TexCoord2fVertex3fSUN func(s, t, x, y, z float32)
// var TexCoord2fVertex3fvSUN func(tc, v *float32)
// var TexCoord2hNV func(s, t GLhalfNV)
// var TexCoord2hvNV func(v *GLhalfNV)
var TexCoord2i func(s, t int)
var TexCoord2iv func(v *int)
var TexCoord2s func(s, t int16)
var TexCoord2sv func(v *int16)
var TexCoord3d func(hs, t, r float64)
var TexCoord3dv func(v *float64)
var TexCoord3f func(hs, t, r float32)
var TexCoord3fv func(v *float32)

// var TexCoord3hNV func(s, t, r GLhalfNV)
// var TexCoord3hvNV func(v *GLhalfNV)
var TexCoord3i func(hs, t, r int)
var TexCoord3iv func(v *int)
var TexCoord3s func(hs, t, r int16)
var TexCoord3sv func(v *int16)
var TexCoord4d func(hs, t, r, q float64)
var TexCoord4dv func(v *float64)
var TexCoord4f func(hs, t, r, q float32)

// var TexCoord4fColor4fNormal3fVertex4fSUN func(s, t, p, q, r, g, b, a, nx, ny, nz, x, y, z, w float32)
// var TexCoord4fColor4fNormal3fVertex4fvSUN func(tc, c, n, v *float32)
var TexCoord4fv func(v *float32)

// var TexCoord4fVertex4fSUN func(s, t, p, q, x, y, z, w float32)
// var TexCoord4fVertex4fvSUN func(tc, v *float32)
// var TexCoord4hNV func(s, t, r, q GLhalfNV)
// var TexCoord4hvNV func(v *GLhalfNV)
var TexCoord4i func(hs, t, r, q int)
var TexCoord4iv func(v *int)
var TexCoord4s func(hs, t, r, q int16)
var TexCoord4sv func(v *int16)

// var TexCoordFormatNV func(size int, type_ Enum, stride Sizei)
var TexCoordP1ui func(type_ Enum, coords uint)
var TexCoordP1uiv func(type_ Enum, coords *uint)
var TexCoordP2ui func(type_ Enum, coords uint)
var TexCoordP2uiv func(type_ Enum, coords *uint)
var TexCoordP3ui func(type_ Enum, coords uint)
var TexCoordP3uiv func(type_ Enum, coords *uint)
var TexCoordP4ui func(type_ Enum, coords uint)
var TexCoordP4uiv func(type_ Enum, coords *uint)
var TexCoordPointer func(size int, type_ Enum, stride Sizei, ptr *GLvoid)
var TexCoordPointerEXT func(size int, type_ Enum, stride, count Sizei, pointer *GLvoid)

// var TexCoordPointerListIBM func(size int, type_ Enum, stride int, pointer **GLvoid, ptrstride int)
// var TexCoordPointervINTEL func(size int, type_ Enum, pointer **GLvoid)
var TexEnvf func(target, pname Enum, param float32)
var TexEnvfv func(target, pname Enum, params *float32)
var TexEnvi func(target, pname Enum, param int)
var TexEnviv func(target, pname Enum, params *int)

// var TexFilterFuncSGIS func(target, filter Enum, n Sizei, weights *float32)
var TexGend func(coord, pname Enum, param float64)
var TexGendv func(coord, pname Enum, params *float64)
var TexGenf func(coord, pname Enum, param float32)
var TexGenfv func(coord, pname Enum, params *float32)
var TexGeni func(coord, pname Enum, param int)
var TexGeniv func(coord, pname Enum, params *int)
var TexImage1D func(target Enum, level, internalFormat int, width Sizei, border int, format, type_ Enum, pixels *GLvoid)
var TexImage2D func(target Enum, level, internalFormat int, width, height Sizei, border int, format, type_ Enum, pixels *GLvoid)

// var TexImage2DMultisample func(target Enum, samples Sizei, internalformat int, width, height Sizei, fixedsamplelocations GLboolean)
// var TexImage2DMultisampleCoverageNV func(target Enum, coverageSamples, colorSamples Sizei, internalFormat int, width, height Sizei, fixedSampleLocations GLboolean)
var TexImage3D func(target Enum, level, internalFormat int, width, height, depth Sizei, border int, format, type_ Enum, pixels *GLvoid)
var TexImage3DEXT func(target Enum, level int, internalformat Enum, width, height, depth Sizei, border int, format, type_ Enum, pixels *GLvoid)

// var TexImage3DMultisample func(target Enum, samples Sizei, internalformat int, width, height, depth Sizei, fixedsamplelocations GLboolean)
// var TexImage3DMultisampleCoverageNV func(target Enum, coverageSamples, colorSamples Sizei, internalFormat int, width, height, depth Sizei, fixedSampleLocations GLboolean)
// var TexImage4DSGIS func(target Enum, level int, internalformat Enum, width, height, depth, size4d Sizei, border int, format, type_ Enum, pixels *GLvoid)
var TexParameterf func(target, pname Enum, param float32)
var TexParameterfv func(target, pname Enum, params *float32)
var TexParameteri func(target, pname Enum, param int)
var TexParameterIiv func(target, pname Enum, params *int)
var TexParameterIivEXT func(target, pname Enum, params *int)
var TexParameterIuiv func(target, pname Enum, params *uint)
var TexParameterIuivEXT func(target, pname Enum, params *uint)
var TexParameteriv func(target, pname Enum, params *int)

// var TexRenderbufferNV func(target Enum, renderbuffer uint)
var TexStorage1D func(target Enum, levels Sizei, internalformat Enum, width Sizei)
var TexStorage2D func(target Enum, levels Sizei, internalformat Enum, width, height Sizei)
var TexStorage3D func(target Enum, levels Sizei, internalformat Enum, width, height, depth Sizei)
var TexSubImage1D func(target Enum, level, xoffset int, width Sizei, format, type_ Enum, pixels *GLvoid)
var TexSubImage1DEXT func(target Enum, level, xoffset int, width Sizei, format, type_ Enum, pixels *GLvoid)
var TexSubImage2D func(target Enum, hlevel, xoffset, yoffset int, width, height Sizei, format, type_ Enum, pixels *GLvoid)
var TexSubImage2DEXT func(target Enum, level, xoffset, yoffset int, width, height Sizei, format, type_ Enum, pixels *GLvoid)
var TexSubImage3D func(target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format, type_ Enum, pixels *GLvoid)
var TexSubImage3DEXT func(target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format, type_ Enum, pixels *GLvoid)

// var TexSubImage4DSGIS func(target Enum, level, xoffset, yoffset, zoffset, woffset int, width, height, depth, size4d Sizei, format, type_ Enum, pixels *GLvoid)
var TextureBarrierNV func()

// var TextureBufferEXT func(texture uint, target, internalformat Enum, buffer uint)
// var TextureColorMaskSGIS func(red, green, blue, alpha GLboolean)
var TextureImage1DEXT func(texture uint, target Enum, level int, internalformat Enum, width Sizei, border int, format, type_ Enum, pixels *GLvoid)
var TextureImage2DEXT func(texture uint, target Enum, level int, internalformat Enum, width, height Sizei, border int, format, type_ Enum, pixels *GLvoid)

// var TextureImage2DMultisampleCoverageNV func(texture uint, target Enum, coverageSamples, colorSamples Sizei, internalFormat int, width, height Sizei, fixedSampleLocations GLboolean)
// var TextureImage2DMultisampleNV func(texture uint, target Enum, samples Sizei, internalFormat int, width, height Sizei, fixedSampleLocations GLboolean)
var TextureImage3DEXT func(texture uint, target Enum, level int, internalformat Enum, width, height, depth Sizei, border int, format, type_ Enum, pixels *GLvoid)

// var TextureImage3DMultisampleCoverageNV func(texture uint, target Enum, coverageSamples, colorSamples Sizei, internalFormat int, width, height, depth Sizei, fixedSampleLocations GLboolean)
// var TextureImage3DMultisampleNV func(texture uint, target Enum, samples Sizei, internalFormat int, width, height, depth Sizei, fixedSampleLocations GLboolean)
// var TextureLightEXT func(pname Enum)
// var TextureMaterialEXT func(face, mode Enum)
// var TextureNormalEXT func(mode Enum)
// var TextureParameterfEXT func(texture uint, target, pname Enum, param float32)
// var TextureParameterfvEXT func(texture uint, target, pname Enum, params *float32)
// var TextureParameteriEXT func(texture uint, target, pname Enum, param int)
// var TextureParameterIivEXT func(texture uint, target, pname Enum, params *int)
// var TextureParameterIuivEXT func(texture uint, target, pname Enum, params *uint)
// var TextureParameterivEXT func(texture uint, target, pname Enum, params *int)
// var TextureRangeAPPLE func(target Enum, length Sizei, pointer *GLvoid)
// var TextureRenderbufferEXT func(texture uint, target Enum, renderbuffer uint)
var TextureStorage1DEXT func(texture uint, target Enum, levels Sizei, internalformat Enum, width Sizei)
var TextureStorage2DEXT func(texture uint, target Enum, levels Sizei, internalformat Enum, width, height Sizei)
var TextureStorage3DEXT func(texture uint, target Enum, levels Sizei, internalformat Enum, width, height, depth Sizei)

// var TextureSubImage1DEXT func(texture uint, target Enum, level, xoffset int, width Sizei, format, type_ Enum, pixels *GLvoid)
// var TextureSubImage2DEXT func(texture uint, target Enum, level, xoffset, yoffset int, width, height Sizei, format, type_ Enum, pixels *GLvoid)
// var TextureSubImage3DEXT func(texture uint, target Enum, level, xoffset, yoffset, zoffset int, width, height, depth Sizei, format, type_ Enum, pixels *GLvoid)
var TrackMatrixNV func(target Enum, address uint, matrix, transform Enum)

// var TransformFeedbackAttribsNV func(count uint, attribs *int, bufferMode Enum)
// var TransformFeedbackStreamAttribsNV func(count Sizei, attribs *int, nbuffers Sizei, bufstreams *int, bufferMode Enum)
var TransformFeedbackVaryings func(program uint, count Sizei, varyings []string, bufferMode Enum)
var TransformFeedbackVaryingsEXT func(program uint, count Sizei, varyings []string, bufferMode Enum)

// var TransformFeedbackVaryingsNV func(program uint, count Sizei, locations *int, bufferMode Enum)
var Translated func(hx, y, z float64)
var Translatef func(hx, y, z float32)

// var Uniform1d func(location int, x float64)
// var Uniform1dv func(location int, count Sizei, value *float64)
var Uniform1f func(location int, v0 float32)
var Uniform1fARB func(location int, v0 float32)
var Uniform1fv func(location int, count Sizei, value *float32)
var Uniform1fvARB func(location int, count Sizei, value *float32)
var Uniform1i func(location, v0 int)

// var Uniform1i64NV func(location, x int64EXT)
// var Uniform1i64vNV func(location int, count Sizei, value *int64EXT)
var Uniform1iARB func(location, v0 int)
var Uniform1iv func(location int, count Sizei, value *int)
var Uniform1ivARB func(location int, count Sizei, value *int)
var Uniform1ui func(location int, v0 uint)

// var Uniform1ui64NV func(location int, x uint64EXT)
// var Uniform1ui64vNV func(location int, count Sizei, value *uint64EXT)
var Uniform1uiEXT func(location int, v0 uint)
var Uniform1uiv func(location int, count Sizei, value *uint)
var Uniform1uivEXT func(location int, count Sizei, value *uint)

// var Uniform2d func(location int, x, y float64)
// var Uniform2dv func(location int, count Sizei, value *float64)
var Uniform2f func(location int, v0, v1 float32)
var Uniform2fARB func(location int, v0, v1 float32)
var Uniform2fv func(location int, count Sizei, value *float32)
var Uniform2fvARB func(location int, count Sizei, value *float32)
var Uniform2i func(location, v0, v1 int)

// var Uniform2i64NV func(location, x, y int64EXT)
// var Uniform2i64vNV func(location int, count Sizei, value *int64EXT)
var Uniform2iARB func(location, v0, v1 int)
var Uniform2iv func(location int, count Sizei, value *int)
var Uniform2ivARB func(location int, count Sizei, value *int)
var Uniform2ui func(location int, v0, v1 uint)

// var Uniform2ui64NV func(location int, x, y uint64EXT)
// var Uniform2ui64vNV func(location int, count Sizei, value *uint64EXT)
var Uniform2uiEXT func(location int, v0, v1 uint)
var Uniform2uiv func(location int, count Sizei, value *uint)
var Uniform2uivEXT func(location int, count Sizei, value *uint)

// var Uniform3d func(location int, x, y, z float64)
// var Uniform3dv func(location int, count Sizei, value *float64)
var Uniform3f func(location int, v0, v1, v2 float32)
var Uniform3fARB func(location int, v0, v1, v2 float32)
var Uniform3fv func(location int, count Sizei, value *float32)
var Uniform3fvARB func(location int, count Sizei, value *float32)
var Uniform3i func(location, v0, v1, v2 int)

// var Uniform3i64NV func(location, x, y, z int64EXT)
// var Uniform3i64vNV func(location int, count Sizei, value *int64EXT)
var Uniform3iARB func(location, v0, v1, v2 int)
var Uniform3iv func(location int, count Sizei, value *int)
var Uniform3ivARB func(location int, count Sizei, value *int)
var Uniform3ui func(location int, v0, v1, v2 uint)

// var Uniform3ui64NV func(location int, x, y, z uint64EXT)
// var Uniform3ui64vNV func(location int, count Sizei, value *uint64EXT)
var Uniform3uiEXT func(location int, v0, v1, v2 uint)
var Uniform3uiv func(location int, count Sizei, value *uint)
var Uniform3uivEXT func(location int, count Sizei, value *uint)

// var Uniform4d func(location int, x, y, z, w float64)
// var Uniform4dv func(location int, count Sizei, value *float64)
var Uniform4f func(location int, v0, v1, v2, v3 float32)
var Uniform4fARB func(location int, v0, v1, v2, v3 float32)
var Uniform4fv func(location int, count Sizei, value *float32)
var Uniform4fvARB func(location int, count Sizei, value *float32)
var Uniform4i func(location, v0, v1, v2, v3 int)

// var Uniform4i64NV func(location, x, y, z, w int64EXT)
// var Uniform4i64vNV func(location int, count Sizei, value *int64EXT)
var Uniform4iARB func(location, v0, v1, v2, v3 int)
var Uniform4iv func(location int, count Sizei, value *int)
var Uniform4ivARB func(location int, count Sizei, value *int)
var Uniform4ui func(location int, v0, v1, v2, v3 uint)

// var Uniform4ui64NV func(location int, x, y, z, w uint64EXT)
// var Uniform4ui64vNV func(location int, count Sizei, value *uint64EXT)
var Uniform4uiEXT func(location int, v0, v1, v2, v3 uint)
var Uniform4uiv func(location int, count Sizei, value *uint)
var Uniform4uivEXT func(location int, count Sizei, value *uint)

// var UniformBlockBinding func(program, uniformBlockIndex, uniformBlockBinding uint)
// var UniformBufferEXT func(program uint, location int, buffer uint)
// var UniformMatrix2dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix2fv func(location int, count Sizei, transpose GLboolean, value *float32)
var UniformMatrix2fvARB func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix2x3dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix2x3fv func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix2x4dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix2x4fv func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix3dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix3fv func(location int, count Sizei, transpose GLboolean, value *float32)
var UniformMatrix3fvARB func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix3x2dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix3x2fv func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix3x4dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix3x4fv func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix4dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix4fv func(location int, count Sizei, transpose GLboolean, value *float32)
var UniformMatrix4fvARB func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix4x2dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix4x2fv func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformMatrix4x3dv func(location int, count Sizei, transpose GLboolean, value *float64)
var UniformMatrix4x3fv func(location int, count Sizei, transpose GLboolean, value *float32)

// var UniformSubroutinesuiv func(shadertype Enum, count Sizei, indices *uint)
// var Uniformui64NV func(location int, value uint64EXT)
// var Uniformui64vNV func(location int, count Sizei, value *uint64EXT)
var UnlockArraysEXT func()
var UnmapBuffer func(target Enum) GLboolean
var UnmapBufferARB func(target Enum) GLboolean

// var UnmapNamedBufferEXT func(buffer uint) GLboolean
// var UnmapObjectBufferATI func(buffer uint)
// var UpdateObjectBufferATI func(buffer, offset uint, size Sizei, pointer *GLvoid, preserve Enum)
var UseProgram func(program uint)
var UseProgramObjectARB func(programObj HandleARB)

// var UseProgramStages func(pipeline uint, stages GLbitfield, program uint)
var UseShaderProgramEXT func(type_ Enum, program uint)
var ValidateProgram func(program uint)
var ValidateProgramARB func(programObj HandleARB)

// var ValidateProgramPipeline func(pipeline uint)
// var VariantArrayObjectATI func(id uint, type_ Enum, stride Sizei, buffer, offset uint)
// var VariantbvEXT func(id uint, addr *GLbyte)
// var VariantdvEXT func(id uint, addr *float64)
// var VariantfvEXT func(id uint, addr *float32)
// var VariantivEXT func(id uint, addr *int)
// var VariantPointerEXT func(id uint, type_ Enum, stride uint, addr *GLvoid)
// var VariantsvEXT func(id uint, addr *int16)
// var VariantubvEXT func(id uint, addr *uint8)
// var VariantuivEXT func(id uint, addr *uint)
// var VariantusvEXT func(id uint, addr *uint16)
// var VDPAUFiniNV func()
// var VDPAUGetSurfaceivNV func(surface GLvdpauSurfaceNV, pname Enum, bufSize Sizei, length *Sizei, values *int)
// var VDPAUInitNV func(vdpDevice, getProcAddress *GLvoid)
// var VDPAUIsSurfaceNV func(surface GLvdpauSurfaceNV)
// var VDPAUMapSurfacesNV func(numSurfaces Sizei, surfaces *GLvdpauSurfaceNV)
// var VDPAURegisterOutputSurfaceNV func(vdpSurface *GLvoid, target Enum, numTextureNames Sizei, textureNames *uint) GLvdpauSurfaceNV
// var VDPAURegisterVideoSurfaceNV func(vdpSurface *GLvoid, target Enum, numTextureNames Sizei, textureNames *uint) GLvdpauSurfaceNV
// var VDPAUSurfaceAccessNV func(surface GLvdpauSurfaceNV, access Enum)
// var VDPAUUnmapSurfacesNV func(numSurface Sizei, surfaces *GLvdpauSurfaceNV)
// var VDPAUUnregisterSurfaceNV func(surface GLvdpauSurfaceNV)
var Vertex2d func(x, y float64)
var Vertex2dv func(v *float64)
var Vertex2f func(x, y float32)
var Vertex2fv func(v *float32)

// var Vertex2hNV func(x, y GLhalfNV)
// var Vertex2hvNV func(v *GLhalfNV)
var Vertex2i func(x, y int)
var Vertex2iv func(v *int)
var Vertex2s func(x, y int16)
var Vertex2sv func(v *int16)
var Vertex3d func(hx, y, z float64)
var Vertex3dv func(v *float64)
var Vertex3f func(hx, y, z float32)
var Vertex3fv func(v *float32)

// var Vertex3hNV func(x, y, z GLhalfNV)
// var Vertex3hvNV func(v *GLhalfNV)
var Vertex3i func(hx, y, z int)
var Vertex3iv func(v *int)
var Vertex3s func(hx, y, z int16)
var Vertex3sv func(v *int16)
var Vertex4d func(hx, y, z, w float64)
var Vertex4dv func(v *float64)
var Vertex4f func(hx, y, z, w float32)
var Vertex4fv func(v *float32)

// var Vertex4hNV func(x, y, z, w GLhalfNV)
// var Vertex4hvNV func(v *GLhalfNV)
var Vertex4i func(hx, y, z, w int)
var Vertex4iv func(v *int)
var Vertex4s func(hx, y, z, w int16)
var Vertex4sv func(v *int16)

// var VertexArrayParameteriAPPLE func(pname Enum, param int)
// var VertexArrayRangeAPPLE func(length Sizei, pointer *GLvoid)
var VertexArrayRangeNV func(length Sizei, pointer *GLvoid)

// var VertexArrayVertexAttribLOffsetEXT func(vaobj, buffer, index uint, size int, type_ Enum, stride Sizei, offset Intptr)
var VertexAttrib1d func(index uint, x float64)
var VertexAttrib1dARB func(index uint, x float64)
var VertexAttrib1dNV func(index uint, x float64)
var VertexAttrib1dv func(index uint, v *float64)
var VertexAttrib1dvARB func(index uint, v *float64)
var VertexAttrib1dvNV func(index uint, v *float64)
var VertexAttrib1f func(index uint, x float32)
var VertexAttrib1fARB func(index uint, x float32)
var VertexAttrib1fNV func(index uint, x float32)
var VertexAttrib1fv func(index uint, v *float32)
var VertexAttrib1fvARB func(index uint, v *float32)
var VertexAttrib1fvNV func(index uint, v *float32)

// var VertexAttrib1hNV func(index uint, x GLhalfNV)
// var VertexAttrib1hvNV func(index uint, v *GLhalfNV)
var VertexAttrib1s func(index uint, x int16)
var VertexAttrib1sARB func(index uint, x int16)
var VertexAttrib1sNV func(index uint, x int16)
var VertexAttrib1sv func(index uint, v *int16)
var VertexAttrib1svARB func(index uint, v *int16)
var VertexAttrib1svNV func(index uint, v *int16)
var VertexAttrib2d func(index uint, x, y float64)
var VertexAttrib2dARB func(index uint, x, y float64)
var VertexAttrib2dNV func(index uint, x, y float64)
var VertexAttrib2dv func(index uint, v *float64)
var VertexAttrib2dvARB func(index uint, v *float64)
var VertexAttrib2dvNV func(index uint, v *float64)
var VertexAttrib2f func(index uint, x, y float32)
var VertexAttrib2fARB func(index uint, x, y float32)
var VertexAttrib2fNV func(index uint, x, y float32)
var VertexAttrib2fv func(index uint, v *float32)
var VertexAttrib2fvARB func(index uint, v *float32)
var VertexAttrib2fvNV func(index uint, v *float32)

// var VertexAttrib2hNV func(index uint, x, y GLhalfNV)
// var VertexAttrib2hvNV func(index uint, v *GLhalfNV)
var VertexAttrib2s func(index uint, x, y int16)
var VertexAttrib2sARB func(index uint, x, y int16)
var VertexAttrib2sNV func(index uint, x, y int16)
var VertexAttrib2sv func(index uint, v *int16)
var VertexAttrib2svARB func(index uint, v *int16)
var VertexAttrib2svNV func(index uint, v *int16)
var VertexAttrib3d func(index uint, x, y, z float64)
var VertexAttrib3dARB func(index uint, x, y, z float64)
var VertexAttrib3dNV func(index uint, x, y, z float64)
var VertexAttrib3dv func(index uint, v *float64)
var VertexAttrib3dvARB func(index uint, v *float64)
var VertexAttrib3dvNV func(index uint, v *float64)
var VertexAttrib3f func(index uint, x, y, z float32)
var VertexAttrib3fARB func(index uint, x, y, z float32)
var VertexAttrib3fNV func(index uint, x, y, z float32)
var VertexAttrib3fv func(index uint, v *float32)
var VertexAttrib3fvARB func(index uint, v *float32)
var VertexAttrib3fvNV func(index uint, v *float32)

// var VertexAttrib3hNV func(index uint, x, y, z GLhalfNV)
// var VertexAttrib3hvNV func(index uint, v *GLhalfNV)
var VertexAttrib3s func(index uint, x, y, z int16)
var VertexAttrib3sARB func(index uint, x, y, z int16)
var VertexAttrib3sNV func(index uint, x, y, z int16)
var VertexAttrib3sv func(index uint, v *int16)
var VertexAttrib3svARB func(index uint, v *int16)
var VertexAttrib3svNV func(index uint, v *int16)
var VertexAttrib4bv func(index uint, v *GLbyte)
var VertexAttrib4bvARB func(index uint, v *GLbyte)
var VertexAttrib4d func(index uint, x, y, z, w float64)
var VertexAttrib4dARB func(index uint, x, y, z, w float64)
var VertexAttrib4dNV func(index uint, x, y, z, w float64)
var VertexAttrib4dv func(index uint, v *float64)
var VertexAttrib4dvARB func(index uint, v *float64)
var VertexAttrib4dvNV func(index uint, v *float64)
var VertexAttrib4f func(index uint, x, y, z, w float32)
var VertexAttrib4fARB func(index uint, x, y, z, w float32)
var VertexAttrib4fNV func(index uint, x, y, z, w float32)
var VertexAttrib4fv func(index uint, v *float32)
var VertexAttrib4fvARB func(index uint, v *float32)
var VertexAttrib4fvNV func(index uint, v *float32)

// var VertexAttrib4hNV func(index uint, x, y, z, w GLhalfNV)
// var VertexAttrib4hvNV func(index uint, v *GLhalfNV)
var VertexAttrib4iv func(index uint, v *int)
var VertexAttrib4ivARB func(index uint, v *int)
var VertexAttrib4Nbv func(index uint, v *GLbyte)
var VertexAttrib4NbvARB func(index uint, v *GLbyte)
var VertexAttrib4Niv func(index uint, v *int)
var VertexAttrib4NivARB func(index uint, v *int)
var VertexAttrib4Nsv func(index uint, v *int16)
var VertexAttrib4NsvARB func(index uint, v *int16)
var VertexAttrib4Nub func(index uint, x, y, z, w uint8)
var VertexAttrib4NubARB func(index uint, x, y, z, w uint8)
var VertexAttrib4Nubv func(index uint, v *uint8)
var VertexAttrib4NubvARB func(index uint, v *uint8)
var VertexAttrib4Nuiv func(index uint, v *uint)
var VertexAttrib4NuivARB func(index uint, v *uint)
var VertexAttrib4Nusv func(index uint, v *uint16)
var VertexAttrib4NusvARB func(index uint, v *uint16)
var VertexAttrib4s func(index uint, x, y, z, w int16)
var VertexAttrib4sARB func(index uint, x, y, z, w int16)
var VertexAttrib4sNV func(index uint, x, y, z, w int16)
var VertexAttrib4sv func(index uint, v *int16)
var VertexAttrib4svARB func(index uint, v *int16)
var VertexAttrib4svNV func(index uint, v *int16)
var VertexAttrib4ubNV func(index uint, x, y, z, w uint8)
var VertexAttrib4ubv func(index uint, v *uint8)
var VertexAttrib4ubvARB func(index uint, v *uint8)
var VertexAttrib4ubvNV func(index uint, v *uint8)
var VertexAttrib4uiv func(index uint, v *uint)
var VertexAttrib4uivARB func(index uint, v *uint)
var VertexAttrib4usv func(index uint, v *uint16)
var VertexAttrib4usvARB func(index uint, v *uint16)

// var VertexAttribArrayObjectATI func(index uint, size int, type_ Enum, normalized GLboolean, stride Sizei, buffer, offset uint)
var VertexAttribDivisor func(index, divisor uint)
var VertexAttribDivisorARB func(index, divisor uint)

// var VertexAttribFormatNV func(index uint, size int, type_ Enum, normalized GLboolean, stride Sizei)
var VertexAttribI1i func(index uint, x int)
var VertexAttribI1iEXT func(index uint, x int)
var VertexAttribI1iv func(index uint, v *int)
var VertexAttribI1ivEXT func(index uint, v *int)
var VertexAttribI1ui func(index, x uint)
var VertexAttribI1uiEXT func(index, x uint)
var VertexAttribI1uiv func(index uint, v *uint)
var VertexAttribI1uivEXT func(index uint, v *uint)
var VertexAttribI2i func(index uint, x, y int)
var VertexAttribI2iEXT func(index uint, x, y int)
var VertexAttribI2iv func(index uint, v *int)
var VertexAttribI2ivEXT func(index uint, v *int)
var VertexAttribI2ui func(index, x, y uint)
var VertexAttribI2uiEXT func(index, x, y uint)
var VertexAttribI2uiv func(index uint, v *uint)
var VertexAttribI2uivEXT func(index uint, v *uint)
var VertexAttribI3i func(index uint, x, y, z int)
var VertexAttribI3iEXT func(index uint, x, y, z int)
var VertexAttribI3iv func(index uint, v *int)
var VertexAttribI3ivEXT func(index uint, v *int)
var VertexAttribI3ui func(index, x, y, z uint)
var VertexAttribI3uiEXT func(index, x, y, z uint)
var VertexAttribI3uiv func(index uint, v *uint)
var VertexAttribI3uivEXT func(index uint, v *uint)
var VertexAttribI4bv func(index uint, v *GLbyte)
var VertexAttribI4bvEXT func(index uint, v *GLbyte)
var VertexAttribI4i func(index uint, x, y, z, w int)
var VertexAttribI4iEXT func(index uint, x, y, z, w int)
var VertexAttribI4iv func(index uint, v *int)
var VertexAttribI4ivEXT func(index uint, v *int)
var VertexAttribI4sv func(index uint, v *int16)
var VertexAttribI4svEXT func(index uint, v *int16)
var VertexAttribI4ubv func(index uint, v *uint8)
var VertexAttribI4ubvEXT func(index uint, v *uint8)
var VertexAttribI4ui func(index, x, y, z, w uint)
var VertexAttribI4uiEXT func(index, x, y, z, w uint)
var VertexAttribI4uiv func(index uint, v *uint)
var VertexAttribI4uivEXT func(index uint, v *uint)
var VertexAttribI4usv func(index uint, v *uint16)
var VertexAttribI4usvEXT func(index uint, v *uint16)

// var VertexAttribIFormatNV func(index uint, size int, type_ Enum, stride Sizei)
var VertexAttribIPointer func(index uint, size int, type_ Enum, stride Sizei, pointer *GLvoid)
var VertexAttribIPointerEXT func(index uint, size int, type_ Enum, stride Sizei, pointer *GLvoid)

// var VertexAttribL1d func(index uint, x float64)
// var VertexAttribL1dEXT func(index uint, x float64)
// var VertexAttribL1dv func(index uint, v *float64)
// var VertexAttribL1dvEXT func(index uint, v *float64)
// var VertexAttribL1i64NV func(index uint, x int64EXT)
// var VertexAttribL1i64vNV func(index uint, v *int64EXT)
// var VertexAttribL1ui64NV func(index, x uint64EXT)
// var VertexAttribL1ui64vNV func(index uint, v *uint64EXT)
// var VertexAttribL2d func(index uint, x, y float64)
// var VertexAttribL2dEXT func(index uint, x, y float64)
// var VertexAttribL2dv func(index uint, v *float64)
// var VertexAttribL2dvEXT func(index uint, v *float64)
// var VertexAttribL2i64NV func(index uint, x, y int64EXT)
// var VertexAttribL2i64vNV func(index uint, v *int64EXT)
// var VertexAttribL2ui64NV func(index, x, y uint64EXT)
// var VertexAttribL2ui64vNV func(index uint, v *uint64EXT)
// var VertexAttribL3d func(index uint, x, y, z float64)
// var VertexAttribL3dEXT func(index uint, x, y, z float64)
// var VertexAttribL3dv func(index uint, v *float64)
// var VertexAttribL3dvEXT func(index uint, v *float64)
// var VertexAttribL3i64NV func(index uint, x, y, z int64EXT)
// var VertexAttribL3i64vNV func(index uint, v *int64EXT)
// var VertexAttribL3ui64NV func(index, x, y, z uint64EXT)
// var VertexAttribL3ui64vNV func(index uint, v *uint64EXT)
// var VertexAttribL4d func(index uint, x, y, z, w float64)
// var VertexAttribL4dEXT func(index uint, x, y, z, w float64)
// var VertexAttribL4dv func(index uint, v *float64)
// var VertexAttribL4dvEXT func(index uint, v *float64)
// var VertexAttribL4i64NV func(index uint, x, y, z, w int64EXT)
// var VertexAttribL4i64vNV func(index uint, v *int64EXT)
// var VertexAttribL4ui64NV func(index, x, y, z, w uint64EXT)
// var VertexAttribL4ui64vNV func(index uint, v *uint64EXT)
// var VertexAttribLFormatNV func(index uint, size int, type_ Enum, stride Sizei)
// var VertexAttribLPointer func(index uint, size int, type_ Enum, stride Sizei, pointer *GLvoid)
// var VertexAttribLPointerEXT func(index uint, size int, type_ Enum, stride Sizei, pointer *GLvoid)
var VertexAttribP1ui func(index uint, type_ Enum, normalized GLboolean, value uint)
var VertexAttribP1uiv func(index uint, type_ Enum, normalized GLboolean, value *uint)
var VertexAttribP2ui func(index uint, type_ Enum, normalized GLboolean, value uint)
var VertexAttribP2uiv func(index uint, type_ Enum, normalized GLboolean, value *uint)
var VertexAttribP3ui func(index uint, type_ Enum, normalized GLboolean, value uint)
var VertexAttribP3uiv func(index uint, type_ Enum, normalized GLboolean, value *uint)
var VertexAttribP4ui func(index uint, type_ Enum, normalized GLboolean, value uint)
var VertexAttribP4uiv func(index uint, type_ Enum, normalized GLboolean, value *uint)
var VertexAttribPointer func(index uint, size int, type_ Enum, normalized GLboolean, stride Sizei, pointer *GLvoid)
var VertexAttribPointerARB func(index uint, size int, type_ Enum, normalized GLboolean, stride Sizei, pointer *GLvoid)
var VertexAttribPointerNV func(index uint, fsize int, type_ Enum, stride Sizei, pointer *GLvoid)
var VertexAttribs1dvNV func(index uint, count Sizei, v *float64)
var VertexAttribs1fvNV func(index uint, count Sizei, v *float32)

// var VertexAttribs1hvNV func(index uint, n Sizei, v *GLhalfNV)
var VertexAttribs1svNV func(index uint, count Sizei, v *int16)
var VertexAttribs2dvNV func(index uint, count Sizei, v *float64)
var VertexAttribs2fvNV func(index uint, count Sizei, v *float32)

// var VertexAttribs2hvNV func(index uint, n Sizei, v *GLhalfNV)
var VertexAttribs2svNV func(index uint, count Sizei, v *int16)
var VertexAttribs3dvNV func(index uint, count Sizei, v *float64)
var VertexAttribs3fvNV func(index uint, count Sizei, v *float32)

// var VertexAttribs3hvNV func(index uint, n Sizei, v *GLhalfNV)
var VertexAttribs3svNV func(index uint, count Sizei, v *int16)
var VertexAttribs4dvNV func(index uint, count Sizei, v *float64)
var VertexAttribs4fvNV func(index uint, count Sizei, v *float32)

// var VertexAttribs4hvNV func(index uint, n Sizei, v *GLhalfNV)
var VertexAttribs4svNV func(index uint, count Sizei, v *int16)
var VertexAttribs4ubvNV func(index uint, count Sizei, v *uint8)

// var VertexBlendARB func(count int)
// var VertexBlendEnvfATI func(pname Enum, param float32)
// var VertexBlendEnviATI func(pname Enum, param int)
// var VertexFormatNV func(size int, type_ Enum, stride Sizei)
var VertexP2ui func(type_ Enum, value uint)
var VertexP2uiv func(type_ Enum, value *uint)
var VertexP3ui func(type_ Enum, value uint)
var VertexP3uiv func(type_ Enum, value *uint)
var VertexP4ui func(type_ Enum, value uint)
var VertexP4uiv func(type_ Enum, value *uint)
var VertexPointer func(size int, type_ Enum, stride Sizei, ptr *GLvoid)
var VertexPointerEXT func(size int, type_ Enum, stride, count Sizei, pointer *GLvoid)

// var VertexPointerListIBM func(size int, type_ Enum, stride int, pointer **GLvoid, ptrstride int)
// var VertexPointervINTEL func(size int, type_ Enum, pointer **GLvoid)
// var VertexStream1dATI func(stream Enum, x float64)
// var VertexStream1dvATI func(stream Enum, coords *float64)
// var VertexStream1fATI func(stream Enum, x float32)
// var VertexStream1fvATI func(stream Enum, coords *float32)
// var VertexStream1iATI func(stream Enum, x int)
// var VertexStream1ivATI func(stream Enum, coords *int)
// var VertexStream1sATI func(stream Enum, x int16)
// var VertexStream1svATI func(stream Enum, coords *int16)
// var VertexStream2dATI func(stream Enum, x, y float64)
// var VertexStream2dvATI func(stream Enum, coords *float64)
// var VertexStream2fATI func(stream Enum, x, y float32)
// var VertexStream2fvATI func(stream Enum, coords *float32)
// var VertexStream2iATI func(stream Enum, x, y int)
// var VertexStream2ivATI func(stream Enum, coords *int)
// var VertexStream2sATI func(stream Enum, x, y int16)
// var VertexStream2svATI func(stream Enum, coords *int16)
// var VertexStream3dATI func(stream Enum, x, y, z float64)
// var VertexStream3dvATI func(stream Enum, coords *float64)
// var VertexStream3fATI func(stream Enum, x, y, z float32)
// var VertexStream3fvATI func(stream Enum, coords *float32)
// var VertexStream3iATI func(stream Enum, x, y, z int)
// var VertexStream3ivATI func(stream Enum, coords *int)
// var VertexStream3sATI func(stream Enum, x, y, z int16)
// var VertexStream3svATI func(stream Enum, coords *int16)
// var VertexStream4dATI func(stream Enum, x, y, z, w float64)
// var VertexStream4dvATI func(stream Enum, coords *float64)
// var VertexStream4fATI func(stream Enum, x, y, z, w float32)
// var VertexStream4fvATI func(stream Enum, coords *float32)
// var VertexStream4iATI func(stream Enum, x, y, z, w int)
// var VertexStream4ivATI func(stream Enum, coords *int)
// var VertexStream4sATI func(stream Enum, x, y, z, w int16)
// var VertexStream4svATI func(stream Enum, coords *int16)
// var VertexWeightfEXT func(weight float32)
// var VertexWeightfvEXT func(weight *float32)
// var VertexWeighthNV func(weight GLhalfNV)
// var VertexWeighthvNV func(weight *GLhalfNV)
// var VertexWeightPointerEXT func(size Sizei, type_ Enum, stride Sizei, pointer *GLvoid)
// var VideoCaptureNV func(video_capture_slot uint, sequence_num, capture_time *uint64EXT) Enum
// var VideoCaptureStreamParameterdvNV func(video_capture_slot, stream uint, pname Enum, params *float64)
// var VideoCaptureStreamParameterfvNV func(video_capture_slot, stream uint, pname Enum, params *float32)
// var VideoCaptureStreamParameterivNV func(video_capture_slot, stream uint, pname Enum, params *int)
var Viewport func(x, y int, width, height Sizei)

// var ViewportArrayv func(first uint, count Sizei, v *float32)
// var ViewportIndexedf func(index uint, x, y, w, h float32)
// var ViewportIndexedfv func(index uint, v *float32)
var WaitSync func(sync Sync, flags GLbitfield, timeout uint64)

// var WeightbvARB func(size int, weights *GLbyte)
// var WeightdvARB func(size int, weights *float64)
// var WeightfvARB func(size int, weights *float32)
// var WeightivARB func(size int, weights *int)
// var WeightPointerARB func(size int, type_ Enum, stride Sizei, pointer *GLvoid)
// var WeightsvARB func(size int, weights *int16)
// var WeightubvARB func(size int, weights *uint8)
// var WeightuivARB func(size int, weights *uint)
// var WeightusvARB func(size int, weights *uint16)
var WindowPos2d func(x, y float64)
var WindowPos2dARB func(x, y float64)
var WindowPos2dMESA func(x, y float64)
var WindowPos2dv func(v *float64)
var WindowPos2dvARB func(v *float64)
var WindowPos2dvMESA func(v *float64)
var WindowPos2f func(x, y float32)
var WindowPos2fARB func(x, y float32)
var WindowPos2fMESA func(x, y float32)
var WindowPos2fv func(v *float32)
var WindowPos2fvARB func(v *float32)
var WindowPos2fvMESA func(v *float32)
var WindowPos2i func(x, y int)
var WindowPos2iARB func(x, y int)
var WindowPos2iMESA func(x, y int)
var WindowPos2iv func(v *int)
var WindowPos2ivARB func(v *int)
var WindowPos2ivMESA func(v *int)
var WindowPos2s func(x, y int16)
var WindowPos2sARB func(x, y int16)
var WindowPos2sMESA func(x, y int16)
var WindowPos2sv func(v *int16)
var WindowPos2svARB func(v *int16)
var WindowPos2svMESA func(v *int16)
var WindowPos3d func(x, y, z float64)
var WindowPos3dARB func(x, y, z float64)
var WindowPos3dMESA func(x, y, z float64)
var WindowPos3dv func(v *float64)
var WindowPos3dvARB func(v *float64)
var WindowPos3dvMESA func(v *float64)
var WindowPos3f func(x, y, z float32)
var WindowPos3fARB func(x, y, z float32)
var WindowPos3fMESA func(x, y, z float32)
var WindowPos3fv func(v *float32)
var WindowPos3fvARB func(v *float32)
var WindowPos3fvMESA func(v *float32)
var WindowPos3i func(x, y, z int)
var WindowPos3iARB func(x, y, z int)
var WindowPos3iMESA func(x, y, z int)
var WindowPos3iv func(v *int)
var WindowPos3ivARB func(v *int)
var WindowPos3ivMESA func(v *int)
var WindowPos3s func(x, y, z int16)
var WindowPos3sARB func(x, y, z int16)
var WindowPos3sMESA func(x, y, z int16)
var WindowPos3sv func(v *int16)
var WindowPos3svARB func(v *int16)
var WindowPos3svMESA func(v *int16)
var WindowPos4dMESA func(x, y, z, w float64)
var WindowPos4dvMESA func(v *float64)
var WindowPos4fMESA func(x, y, z, w float32)
var WindowPos4fvMESA func(v *float32)
var WindowPos4iMESA func(x, y, z, w int)
var WindowPos4ivMESA func(v *int)
var WindowPos4sMESA func(x, y, z, w int16)
var WindowPos4svMESA func(v *int16)

// var WriteMaskEXT func(res, in uint, outX, outY, outZ, outW Enum)

var BeginCurve func(nurb *GLUnurbs)
var BeginPolygon func(tess *GLUtesselator)
var BeginSurface func(nurb *GLUnurbs)
var BeginTrim func(nurb *GLUnurbs)
var Build1DMipmapLevels func(target Enum, internalFormat int, width Sizei, format, type_ Enum, level, base, max int, data *void) int
var Build1DMipmaps func(target Enum, internalFormat int, width Sizei, format, type_ Enum, data *void) int
var Build2DMipmapLevels func(target Enum, internalFormat int, width, height Sizei, format, type_ Enum, level, base, max int, data *void) int
var Build2DMipmaps func(target Enum, internalFormat int, width, height Sizei, format, type_ Enum, data *void) int
var Build3DMipmapLevels func(target Enum, internalFormat int, width, height, depth Sizei, format, type_ Enum, level, base, max int, data *void) int
var Build3DMipmaps func(target Enum, internalFormat int, width, height, depth Sizei, format, type_ Enum, data *void) int
var CheckExtension func(extName, extString *uint8) GLboolean
var Cylinder func(quad *GLUquadric, base, top, height float64, slices, stacks int)
var DeleteNurbsRenderer func(nurb *GLUnurbs)
var DeleteQuadric func(quad *GLUquadric)
var DeleteTess func(tess *GLUtesselator)
var Disk func(quad *GLUquadric, inner, outer float64, slices, loops int)
var EndCurve func(nurb *GLUnurbs)
var EndPolygon func(tess *GLUtesselator)
var EndSurface func(nurb *GLUnurbs)
var EndTrim func(nurb *GLUnurbs)
var ErrorString func(err Enum) *uint8
var GetNurbsProperty func(nurb *GLUnurbs, property Enum, data *float32)

// var GetString func(name Enum) *uint8
var GetTessProperty func(tess *GLUtesselator, which Enum, data *float64)
var LoadSamplingMatrices func(nurb *GLUnurbs, model, perspective *float32, view *int)
var LookAt func(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64)
var NewNurbsRenderer func(void) *GLUnurbs
var NewQuadric func(void) *GLUquadric
var NewTess func(void) *GLUtesselator
var NextContour func(tess *GLUtesselator, type_ Enum)
var NurbsCallback func(nurb *GLUnurbs, which Enum, CallBackFunc Funcptr)
var NurbsCallbackData func(nurb *GLUnurbs, userData *GLvoid)
var NurbsCallbackDataEXT func(nurb *GLUnurbs, userData *GLvoid)
var NurbsCurve func(nurb *GLUnurbs, knotCount int, knots *float32, stride int, control *float32, order int, type_ Enum)
var NurbsProperty func(nurb *GLUnurbs, property Enum, value float32)
var NurbsSurface func(nurb *GLUnurbs, sKnotCount int, sKnots *float32, tKnotCount int, tKnots *float32, sStride, tStride int, control *float32, sOrder, tOrder int, type_ Enum)
var Ortho2D func(left, right, bottom, top float64)
var PartialDisk func(quad *GLUquadric, inner, outer float64, slices, loops int, start, sweep float64)
var Perspective func(fovy, aspect, zNear, zFar float64)
var PickMatrix func(x, y, delX, delY float64, viewport *int)
var Project func(objX, objY, objZ float64, model, proj *float64, view *int, winX, winY, winZ *float64) int
var PwlCurve func(nurb *GLUnurbs, count int, data *float32, stride int, type_ Enum)
var QuadricCallback func(quad *GLUquadric, which Enum, CallBackFunc Funcptr)
var QuadricDrawStyle func(quad *GLUquadric, draw Enum)
var QuadricNormals func(quad *GLUquadric, normal Enum)
var QuadricOrientation func(quad *GLUquadric, orientation Enum)
var QuadricTexture func(quad *GLUquadric, texture GLboolean)
var ScaleImage func(format Enum, wIn, hIn Sizei, typeIn Enum, dataIn *void, wOut, hOut Sizei, typeOut Enum, dataOut *GLvoid) int
var Sphere func(quad *GLUquadric, radius float64, slices, stacks int)
var TessBeginContour func(tess *GLUtesselator)
var TessBeginPolygon func(tess *GLUtesselator, data *GLvoid)
var TessCallback func(tess *GLUtesselator, which Enum, CallBackFunc Funcptr)
var TessEndContour func(tess *GLUtesselator)
var TessEndPolygon func(tess *GLUtesselator)
var TessNormal func(tess *GLUtesselator, valueX, valueY, valueZ float64)
var TessProperty func(tess *GLUtesselator, which Enum, data float64)
var TessVertex func(tess *GLUtesselator, location *float64, data *GLvoid)
var UnProject func(winX, winY, winZ float64, model, proj *float64, view *int, objX, objY, objZ *float64) int
var UnProject4 func(winX, winY, winZ, clipW float64, model, proj *float64, view *int, nearVal, farVal float64, objX, objY, objZ, objW *float64) int

type Enum uint

const (
	Points Enum = iota
	Lines
	LineLoop
	LineStrip
	Triangles
	TriangleStrip
	TriangleFan
	Quads
	QuadStrip
	Polygon
)

const (
	MatrixMode_ = 0x0BA0 //TODO(t):name conflict func
	ModelView   = 0x1700
	Projection  = 0x1701
	Texture     = 0x1702
)

const (
	Lighting               = 0x0B50
	Light0                 = 0x4000
	Light1                 = 0x4001
	Light2                 = 0x4002
	Light3                 = 0x4003
	Light4                 = 0x4004
	Light5                 = 0x4005
	Light6                 = 0x4006
	Light7                 = 0x4007
	SpotExponent           = 0x1205
	SpotCutoff             = 0x1206
	ConstantAttenuation    = 0x1207
	LinearAttenuation      = 0x1208
	QuadraticAttenuation   = 0x1209
	Ambient                = 0x1200
	Diffuse                = 0x1201
	Specular               = 0x1202
	Shininess              = 0x1601
	Emission               = 0x1600
	Position               = 0x1203
	SpotDirection          = 0x1204
	AmbientAndDiffuse      = 0x1602
	ColorIndexes           = 0x1603
	LightModelTwoSide      = 0x0b52
	LightModelLocalViewer  = 0x0b51
	LightModelAmbient      = 0x0b53
	FrontAnd_back          = 0x0408
	SHADEModel             = 0x0B54
	COLORMaterial          = 0x0B57
	ColorMaterialFace      = 0x0B55
	ColorMaterialParameter = 0x0B56
	Mormalize              = 0x0BA1
)

const (
	Point               = 0x1B00
	Line                = 0x1B01
	Fill                = 0x1B02
	Cw                  = 0x0900
	Ccw                 = 0x0901
	Front               = 0x0404
	Back                = 0x0405
	POLYGONMode         = 0x0B40
	PolygonSmooth       = 0x0B41
	POLYGONStipple      = 0x0B42
	EDGEFlag            = 0x0B43
	CULLFace            = 0x0B44
	CullFaceMode        = 0x0B45
	FRONTFace           = 0x0B46
	PolygonOffsetFactor = 0x8038
	PolygonOffsetUnits  = 0x2A00
	PolygonOffsetPoint  = 0x2A01
	PolygonOffsetLine   = 0x2A02
	PolygonOffsetFill   = 0x8037
)

const (
	Never           = 0x0200
	Less            = 0x0201
	Equal           = 0x0202
	LEqual          = 0x0203
	Greater         = 0x0204
	NotEqual        = 0x0205
	GEqual          = 0x0206
	Always          = 0x0207
	DepthTest       = 0x0B71
	DepthBits       = 0x0D56
	DepthClearValue = 0x0B73
	DEPTHFunc       = 0x0B74
	DEPTHRange      = 0x0B70
	DepthWriteMask  = 0x0B72
	DepthComponent  = 0x1902
)

const (
	Compile           = 0x1300
	CompileAndExecute = 0x1301
	LISTBase          = 0x0B32
	LISTIndex         = 0x0B33
	ListMode          = 0x0B30
)
