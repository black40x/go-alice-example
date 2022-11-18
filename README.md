# Alice Skill Example

This is linear Yandex.Dialog Alice skill example with YAML dialog structure written with GO.

### How to use?

Linear dialog load from `dialog.yaml` and have this structure:

```yaml
name: "Your dialog name"
global-errors:
  - text: "Error message said by alice"
dialog:
  - alice:
      text: "Alice text"
      tts: "Alice text TTS"
    user: "user linear answer (request.command) with lower case words only, be careful and test it." 
```

### Deploy by Yandex.Cloud functions

For deploy to Yandex.Cloud functions you can use this easy steps:

- Download this repository as ZIP archive
- Open your Yandex.Cloud [console](https://console.cloud.yandex.ru/) and move to **Cloud Functions**.
- Create new function
- Open "Editor" tab, select "Go 1.19" and uncheck examples checkbox
- Use source codes from ZIP archive
- Set `dialog.Handler` as entry point
- Click "Create version"

Next, you need to create an Alice dialog with these steps:
- Open [Yandex.Dialogs](https://dialogs.yandex.ru/)
- Create dialog with Alice Skill
- Select your cloud function tn the "Backend" field
- Save and publish it.