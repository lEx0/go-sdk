// Copyright (c) 2020 YANDEX LLC.

package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/ai/stt"
	"github.com/yandex-cloud/go-sdk/gen/ai/translate"
	"github.com/yandex-cloud/go-sdk/gen/ai/tts"
	"github.com/yandex-cloud/go-sdk/gen/ai/vision"
)

const (
	AITranslate Endpoint = "ai-translate"
	AIVision    Endpoint = "ai-vision"
	AISTT       Endpoint = "ai-stt"
	AISpeechKit Endpoint = "ai-speechkit"
)

type AI struct {
	sdk *SDK
}

func (m *AI) Translate() *translate.Translate {
	return translate.NewTranslate(m.sdk.getConn(AITranslate))
}

func (m *AI) Vision() *vision.Vision {
	return vision.NewVision(m.sdk.getConn(AIVision))
}

func (m *AI) STT() *stt.STT {
	return stt.NewSTT(m.sdk.getConn(AISTT))
}

func (m *AI) TTS() *tts.TTS {
	return tts.NewTTS(m.sdk.getConn(AISpeechKit))
}
