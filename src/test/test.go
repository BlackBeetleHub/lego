package main

import (
	"encoding/json"
	"strings"
	"io"
	"log"
	"fmt"
	json_response "json"
)

func main() {
	const jsonStream = `
			{
			  "error_msg": "",
			  "count_words": 3489,
			  "show_more": true,
			  "userdict3": [
				{
				  "name": "Сегодня",
				  "count": 109,
				  "data": "1504310400:1504374540",
				  "words": [
					{
					  "word_id": 29913,
					  "word_value": "pen",
					  "word_top": 3,
					  "transcription": "pɛn",
					  "progress_percent": 0,
					  "meeting_count": 0,
					  "training_state": 4095,
					  "training_count": 0,
					  "created_at": 1504356489,
					  "last_updated_at": 1504356489,
					  "user_translates": [
						{
						  "translate_id": 145844,
						  "translate_value": "ручка",
						  "speech_part_id": 0,
						  "translate_votes": 34202,
						  "picture_id": 3221083,
						  "picture_url": "//contentcdn.lingualeo.com/uploads/picture/3221083.png",
						  "pictures": {
							"3465479": "//contentcdn.lingualeo.com/uploads/picture/3465479.png",
							"3463462": "//contentcdn.lingualeo.com/uploads/picture/3463462.png",
							"3413546": "//contentcdn.lingualeo.com/uploads/picture/3413546.png",
							"2907076": "//contentcdn.lingualeo.com/uploads/picture/2907076.png",
							"2984670": "//contentcdn.lingualeo.com/uploads/picture/2984670.png",
							"825661": "//contentcdn.lingualeo.com/uploads/picture/825661.png",
							"863543": "//contentcdn.lingualeo.com/uploads/picture/863543.png",
							"863544": "//contentcdn.lingualeo.com/uploads/picture/863544.png",
							"873634": "//contentcdn.lingualeo.com/uploads/picture/873634.png",
							"705561": "//contentcdn.lingualeo.com/uploads/picture/705561.png",
							"725241": "//contentcdn.lingualeo.com/uploads/picture/725241.png",
							"755877": "//contentcdn.lingualeo.com/uploads/picture/755877.png",
							"671584": "//contentcdn.lingualeo.com/uploads/picture/671584.png",
							"588947": "//contentcdn.lingualeo.com/uploads/picture/588947.png",
							"605651": "//contentcdn.lingualeo.com/uploads/picture/605651.png",
							"528573": "//contentcdn.lingualeo.com/uploads/picture/528573.png",
							"440683": "//contentcdn.lingualeo.com/uploads/picture/440683.png",
							"423631": "//contentcdn.lingualeo.com/uploads/picture/423631.png",
							"428236": "//contentcdn.lingualeo.com/uploads/picture/428236.png"
						  }
						}
					  ],
					  "picture_id": 3221083,
					  "picture_url": "//contentcdn.lingualeo.com/uploads/picture/3221083.png",
					  "pictures": {
						"3465479": "//contentcdn.lingualeo.com/uploads/picture/3465479.png",
						"3463462": "//contentcdn.lingualeo.com/uploads/picture/3463462.png",
						"3413546": "//contentcdn.lingualeo.com/uploads/picture/3413546.png",
						"2907076": "//contentcdn.lingualeo.com/uploads/picture/2907076.png",
						"2984670": "//contentcdn.lingualeo.com/uploads/picture/2984670.png",
						"825661": "//contentcdn.lingualeo.com/uploads/picture/825661.png",
						"863543": "//contentcdn.lingualeo.com/uploads/picture/863543.png",
						"863544": "//contentcdn.lingualeo.com/uploads/picture/863544.png",
						"873634": "//contentcdn.lingualeo.com/uploads/picture/873634.png",
						"705561": "//contentcdn.lingualeo.com/uploads/picture/705561.png",
						"725241": "//contentcdn.lingualeo.com/uploads/picture/725241.png",
						"755877": "//contentcdn.lingualeo.com/uploads/picture/755877.png",
						"671584": "//contentcdn.lingualeo.com/uploads/picture/671584.png",
						"588947": "//contentcdn.lingualeo.com/uploads/picture/588947.png",
						"605651": "//contentcdn.lingualeo.com/uploads/picture/605651.png",
						"528573": "//contentcdn.lingualeo.com/uploads/picture/528573.png",
						"440683": "//contentcdn.lingualeo.com/uploads/picture/440683.png",
						"423631": "//contentcdn.lingualeo.com/uploads/picture/423631.png",
						"428236": "//contentcdn.lingualeo.com/uploads/picture/428236.png"
					  },
					  "note": "",
					  "context_id": 120639,
					  "context": "The pen where they were fed.",
					  "context_content_id": 0,
					  "sound_url": "http://audiocdn.lingualeo.com/v2/1/29913-631152008.mp3",
					  "user_word_value": true,
					  "groups": [
						235
					  ]
					}
				  ]
				}
			  ],
			  "experienceSkills": null,
			  "questData": {
				"task_num": 20,
				"task_num_prev": 0,
				"task_state": 1,
				"task_actions_finished": [
				  [
					1,
					1
				  ]
				],
				"meatballs": 2727,
				"leoClothing": 8
			  },
			  "meatballs": 2727,
			  "notify_count": 0,
			  "utcServerTime": 1504378199,
			  "_hash": "1707.1.3-rc10"
			}`


	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m json_response.LeoDictionaryImpl
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Print( m.User_dict[0].Words[0].Transcription)
	}
}
