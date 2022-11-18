package main

import (
	"encoding/json"
	"math/rand"
)

func Handler(req []byte) ([]byte, error) {
	var reqJson JsonData
	json.Unmarshal(req, &reqJson)

	body := JsonData{
		"version": reqJson["version"],
		"session": reqJson["session"],
	}

	scenario, err := OpenScenario("./dialog.yaml")
	if err != nil {
		body["response"] = JsonData{
			"text":        "Диалог не наиден",
			"end_session": true,
		}
	} else {
		// Start dialog
		if ns, _ := reqJson.GetBool("session.new"); ns == true {
			body["session_state"] = JsonData{
				"value": 0,
			}
			body["response"] = JsonData{
				"text":        scenario.Dialog[0].Alice.Text,
				"tts":         scenario.Dialog[0].Alice.TTS,
				"end_session": false,
			}
		} else {
			currentId, _ := reqJson.GetFloat64("state.session.value")
			command, _ := reqJson.GetString("request.command")
			node := scenario.Dialog[int(currentId)]

			if command == node.User || node.User == "" {
				// ToDo - Extract
				body["session_state"] = JsonData{
					"value": currentId + 1,
				}

				next := scenario.Dialog[int(currentId+1)]

				body["response"] = JsonData{
					"text":        next.Alice.Text,
					"tts":         next.Alice.TTS,
					"end_session": currentId+1 == float64(len(scenario.Dialog)-1),
				}
			} else {
				errorText := "Ответ не верный! Попробуй еще разок..."
				if len(scenario.GlobalErrors) > 0 {
					errorText = scenario.GlobalErrors[rand.Intn(len(scenario.GlobalErrors)-1)].Text
				}

				body["response"] = JsonData{
					"text":        errorText,
					"end_session": false,
				}
				body["session_state"] = JsonData{
					"value": currentId,
				}
			}
		}
	}

	return json.Marshal(body)
}
