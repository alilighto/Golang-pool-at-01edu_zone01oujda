curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq --arg ID_num "$HERO_ID" '.[] | select(.id == ($ID_num|tonumber)) | .connections.relatives' | sed 's/"//g'
