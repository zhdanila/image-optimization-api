// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package image

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	bind "image-optimization-api/pkg/bind"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage(in *jlexer.Lexer, out *UploadImageRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Images":
			if in.IsNull() {
				in.Skip()
				out.Images = nil
			} else {
				in.Delim('[')
				if out.Images == nil {
					if !in.IsDelim(']') {
						out.Images = make([]bind.UploadedFile, 0, 0)
					} else {
						out.Images = []bind.UploadedFile{}
					}
				} else {
					out.Images = (out.Images)[:0]
				}
				for !in.IsDelim(']') {
					var v1 bind.UploadedFile
					easyjsonCce3d1beDecodeImageOptimizationApiPkgBind(in, &v1)
					out.Images = append(out.Images, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage(out *jwriter.Writer, in UploadImageRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Images) != 0 {
		const prefix string = ",\"Images\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.Images {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonCce3d1beEncodeImageOptimizationApiPkgBind(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UploadImageRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UploadImageRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UploadImageRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UploadImageRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage(l, v)
}
func easyjsonCce3d1beDecodeImageOptimizationApiPkgBind(in *jlexer.Lexer, out *bind.UploadedFile) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "filename":
			out.FileName = string(in.String())
		case "ContentType":
			out.ContentType = string(in.String())
		case "Size":
			out.Size = int64(in.Int64())
		case "Src":
			if in.IsNull() {
				in.Skip()
				out.Src = nil
			} else {
				out.Src = in.Bytes()
			}
		case "Tag":
			out.Tag = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiPkgBind(out *jwriter.Writer, in bind.UploadedFile) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FileName != "" {
		const prefix string = ",\"filename\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.FileName))
	}
	if in.ContentType != "" {
		const prefix string = ",\"ContentType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContentType))
	}
	if in.Size != 0 {
		const prefix string = ",\"Size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Size))
	}
	if len(in.Src) != 0 {
		const prefix string = ",\"Src\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Src)
	}
	if in.Tag != "" {
		const prefix string = ",\"Tag\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Tag))
	}
	out.RawByte('}')
}
func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage1(in *jlexer.Lexer, out *ListOriginImageResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "keys":
			if in.IsNull() {
				in.Skip()
				out.Keys = nil
			} else {
				in.Delim('[')
				if out.Keys == nil {
					if !in.IsDelim(']') {
						out.Keys = make([]string, 0, 4)
					} else {
						out.Keys = []string{}
					}
				} else {
					out.Keys = (out.Keys)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Keys = append(out.Keys, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage1(out *jwriter.Writer, in ListOriginImageResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Keys) != 0 {
		const prefix string = ",\"keys\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v8, v9 := range in.Keys {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListOriginImageResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListOriginImageResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListOriginImageResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListOriginImageResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage1(l, v)
}
func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage2(in *jlexer.Lexer, out *ListOriginImageRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage2(out *jwriter.Writer, in ListOriginImageRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListOriginImageRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListOriginImageRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListOriginImageRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListOriginImageRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage2(l, v)
}
func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage3(in *jlexer.Lexer, out *ListImageResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "images":
			if in.IsNull() {
				in.Skip()
				out.Images = nil
			} else {
				in.Delim('[')
				if out.Images == nil {
					if !in.IsDelim(']') {
						out.Images = make([]Info, 0, 2)
					} else {
						out.Images = []Info{}
					}
				} else {
					out.Images = (out.Images)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Info
					(v10).UnmarshalEasyJSON(in)
					out.Images = append(out.Images, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage3(out *jwriter.Writer, in ListImageResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Images) != 0 {
		const prefix string = ",\"images\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v11, v12 := range in.Images {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListImageResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListImageResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListImageResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListImageResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage3(l, v)
}
func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage4(in *jlexer.Lexer, out *ListImageRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage4(out *jwriter.Writer, in ListImageRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListImageRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListImageRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListImageRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListImageRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage4(l, v)
}
func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage5(in *jlexer.Lexer, out *Info) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "key":
			out.Key = string(in.String())
		case "url":
			out.URL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage5(out *jwriter.Writer, in Info) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Key != "" {
		const prefix string = ",\"key\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Key))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Info) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Info) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Info) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Info) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage5(l, v)
}
func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage6(in *jlexer.Lexer, out *GetImageResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "image":
			(out.Image).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage6(out *jwriter.Writer, in GetImageResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"image\":"
		first = false
		out.RawString(prefix[1:])
		(in.Image).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetImageResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetImageResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetImageResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetImageResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage6(l, v)
}
func easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage7(in *jlexer.Lexer, out *GetImageRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ImageID":
			out.ImageID = string(in.String())
		case "CompressionQuality":
			out.CompressionQuality = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage7(out *jwriter.Writer, in GetImageRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ImageID != "" {
		const prefix string = ",\"ImageID\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ImageID))
	}
	if in.CompressionQuality != 0 {
		const prefix string = ",\"CompressionQuality\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CompressionQuality))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetImageRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetImageRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCce3d1beEncodeImageOptimizationApiInternalServiceImage7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetImageRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetImageRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCce3d1beDecodeImageOptimizationApiInternalServiceImage7(l, v)
}
