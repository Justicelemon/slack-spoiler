Slack Spoilers
==============

Slack Spoilers is an API designed for hiding spoilers in Slack messages by allowing users to quickly get an encoded string for their spoiler-filled message and anyone can then copy and paste the encoded message and then decode with a different API call. The encoding/decoding is done using base64 encoding/decoding. This was designed to be used as a Slash Command on Slack. 


Dependencies
-----
- Python 2.7
- Google App Engine SDK for Go

Running Locally
-----
Run `goapp serve app/` in the main directory or just `goapp serve` in the app directory to start the local development server. The app should now be accessible at `http://localhost:8080`, though just typing `http://localhost:8080` without any commands leads to the app responding with "I do not understand your command". The development app server watches for changes in the files, so there's no need to restart `goapp serve` as development is being done on it.

Encoding: http://localhost:8080/?command=/encode&text={TEXT_TO_ENCODE_HERE}

Decoding: http://localhost:8080/?command=/decode&text={ENCODED_TEXT_TO_DECODE_HERE}


Deploying
-----
1. Sign in to [App Engine](https://appengine.google.com/) with yoru Google Account and click "Create Application"
2. Fill in the form for creating your application, make sure to note what you entered for "Application Identifier"
3. Change the value of the `application:` setting in app.yaml from slackspoiler to the Application Identifier set in step 2
4. Upload the app by running `goapp deploy app/` and enter your Google email ID and password

The app should now be available at `http://{app_id}.appspot.com`
