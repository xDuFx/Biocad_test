package models

type DataTcv struct {
	N          int    `tsv:"n"`
	Mqtt       string `tsv:"mqtt"`
	Invid      string `tsv:"invid"`
	Unit_guid  string `tsv:"unit_guimsg_id"`
	Msg_id     string `tsv:"msg_id"`
	Text       string `tsv:"text"`
	Context    string `tsv:"context"`
	Class      string `tsv:"class "`
	Level      int    `tsv:"level"`
	Area       string `tsv:"area"`
	Addr       string `tsv:"addr"`
	Block      string `tsv:"block"`
	TypE       string `tsv:"type"`
	Bit        string `tsv:"bit"`
	Invert_bit string `tsv:"invert_bit"`
}

