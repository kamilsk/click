// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF1a92ff2DecodeGithubComKamilskClickPkgDomain(in *jlexer.Lexer, out *Rule) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "description":
			out.Description = string(in.String())
		case "alias":
			out.AliasID = ID(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tags = append(out.Tags, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "conditions":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Conditions = make(map[string]string)
				} else {
					out.Conditions = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string
					v2 = string(in.String())
					(out.Conditions)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "match":
			out.Match = uint8(in.Uint8())
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
func easyjsonF1a92ff2EncodeGithubComKamilskClickPkgDomain(out *jwriter.Writer, in Rule) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.AliasID != "" {
		const prefix string = ",\"alias\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AliasID))
	}
	if len(in.Tags) != 0 {
		const prefix string = ",\"tags\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Tags {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Conditions {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.String(string(v5Value))
			}
			out.RawByte('}')
		}
	}
	if in.Match != 0 {
		const prefix string = ",\"match\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint8(uint8(in.Match))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Rule) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF1a92ff2EncodeGithubComKamilskClickPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Rule) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF1a92ff2EncodeGithubComKamilskClickPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Rule) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF1a92ff2DecodeGithubComKamilskClickPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Rule) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF1a92ff2DecodeGithubComKamilskClickPkgDomain(l, v)
}
