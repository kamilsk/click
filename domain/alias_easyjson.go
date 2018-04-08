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

func easyjson8922032aDecodeGithubComKamilskClickDomain(in *jlexer.Lexer, out *Alias) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "namespace":
			out.Namespace = string(in.String())
		case "urn":
			out.URN = string(in.String())
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
func easyjson8922032aEncodeGithubComKamilskClickDomain(out *jwriter.Writer, in Alias) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	{
		const prefix string = ",\"urn\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URN))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Alias) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8922032aEncodeGithubComKamilskClickDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Alias) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8922032aEncodeGithubComKamilskClickDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Alias) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8922032aDecodeGithubComKamilskClickDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Alias) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8922032aDecodeGithubComKamilskClickDomain(l, v)
}