package main

// データ構造体の定義。タイトルと本文
type Page struct {
	Title string
	Body  []byte
	// Body は何故stringじゃないのか？→使用するioライブラリで予期されるタイプであるため（？）
}

// Saveメソッド作成から-----
