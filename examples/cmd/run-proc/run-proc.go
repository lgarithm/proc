package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/lgarithm/proc-experimental"
	"github.com/lgarithm/proc-experimental/builtin"
	"github.com/lgarithm/proc-experimental/control"
	"github.com/lgarithm/proc-experimental/execution"
	"github.com/lgarithm/proc-experimental/iostream"
	"github.com/lgarithm/proc-experimental/xterm"
)

func main() {
	flag.Parse()
	parExample()
	seqExample()
	tryExample()
	complexExample()
}

func parExample() {
	p := proc.Proc{
		Prog: `task`,
		Args: []string{`par-example`},
	}
	{
		p := control.Par(
			builtin.Shell(p.CmdCtx(context.TODO())),
			builtin.Shell(p.CmdCtx(context.TODO())),
		)
		w := iostream.NewXTermRedirector(`x`, xterm.Green)
		if r := execution.Run(p, w); r.Err != nil {
			fmt.Printf("failed: %v\n", r.Err)
		}
	}
}

func seqExample() {
	p := proc.Proc{
		Prog: `task`,
		Args: []string{`seq-example`},
	}
	{
		p := control.Seq(
			builtin.Shell(p.CmdCtx(context.TODO())),
			builtin.Shell(p.CmdCtx(context.TODO())),
		)
		w := iostream.NewXTermRedirector(`x`, xterm.Green)
		if r := execution.Run(p, w); r.Err != nil {
			fmt.Printf("failed: %v\n", r.Err)
		}
	}
}

func tryExample() {
	e := errors.New("e")
	var n int
	q := func() execution.P {
		n++
		fmt.Printf("trial #%d\n", n)
		return control.Par(
			builtin.RandomFailure(e, 0.9, rand.New(rand.NewSource(time.Now().UnixNano()))),
			builtin.RandomFailure(e, 0.9, rand.New(rand.NewSource(time.Now().UnixNano()))),
			builtin.RandomFailure(e, 0.9, rand.New(rand.NewSource(time.Now().UnixNano()))),
		)
	}
	p := control.Try(q)
	w := iostream.NewXTermRedirector(`x`, xterm.Green)
	if r := execution.Run(p, w); r.Err != nil {
		fmt.Printf("failed: %v\n", r.Err)
	}
}

const narrative = `narrative
記憶の層に隠れてしまう
夢の色に憧れた君は
強さを装った大人の陰を
踏まないように光を探した
時計の針に押されそうな
負け惜しみなどいらない
誰かの手を掴みそうに
倒れていくなら
痛みと立ち上がっていく涙で描く
埃と空の間を
落ちかけている子供の丘で
逃げた理由を背中に隠してる
未来の声に脅かされそうで
僕を守る君は震えていた
後悔に慣れた独り言
歩き出す音で消した
答えを出せない過去達に
庇われるなら
知らない明日を着て転べばいい
無邪気に咲いた願いと
わがままな希望 幼い顔で笑う
怯える怒りがこの身体を支えてる
言葉ばかりの雨に消されないように
今の灯りを目の前の答えとつなぐ
誰かの手を掴みそうに
倒れていくなら
痛みと立ち上がっていく涙で描く
埃と空の間を
わがままな希望 幼い顔で笑う
怯える怒りがこの身体を支えてる
言葉ばかりの雨に消されないように
今の灯りを目の前の答えとつなぐ
`

func complexExample() {
	var n int
	nt := func() execution.P {
		n++
		prefix := func(x string) string { return fmt.Sprintf("[nt/%d/%s] ", n, x) }
		p1 := control.Term(prefix(`A`), builtin.Echo(narrative))
		p2 := control.Term(prefix(`B`), builtin.Echo(narrative))
		p3 := control.Term(prefix(`C`), builtin.Echo(narrative))
		return control.Seq(p1, p2, p3)
	}
	p := control.Par(nt(), nt(), nt())
	w := iostream.NewTerminalRedirector(``)
	r := execution.Run(p, w)
	r.Unwrap()
}
