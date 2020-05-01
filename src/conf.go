// 設定ファイルの読み出し

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	confFileName = "conf.json"
)

// Config は設定情報の型
type Config struct {
	Width     int   `json:"width"`
	Height    int   `json:"height"`
	InitCells []int `json:"init_cells"`
	Debug     bool  `json:"debug"`
}

// LoadConfig はファイルに保存された JSON オブジェクトを読み出す関数
func LoadConfig() (conf Config, err error) {

	// バイト列読み出し
	var bytes []byte
	bytes, err = ioutil.ReadFile(resolvConfFile())
	if err != nil {
		return
	}

	// json 形式のデコード
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return
	}

	// 必要ならば、ここで conf の格納値をチェックする。
	if conf.Width != 9 {
		conf.Width = 9
	}
	if conf.Height != 9 {
		conf.Height = 9
	}
	if conf.InitCells == nil || len(conf.InitCells) != conf.Width*conf.Height {
		conf.InitCells = make([]int, conf.Width*conf.Height)
	}
	return
}

// resolvConfFile は設定ファイルのパスを特定する関数。
// 実行ファイルと同じディレクトリ配下の設定ファイルのパスとする。
func resolvConfFile() string {
	// 実行ファイルのパスを特定
	exe, err := os.Executable()
	if err == nil {
		// 実行ファイルのあるディレクトリ配下の設定ファイルのパス
		return filepath.Dir(exe) + "/" + confFileName
	}

	// 実行カレントディレクトリ配下の設定ファイルのパス
	return confFileName
}
