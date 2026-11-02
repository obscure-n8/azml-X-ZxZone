
<p align="center">
    <i>A high-performance, ultra-lightweight Telegram Leech Bot engineered in pure Go. Built on top of the Gotd MTProto framework, featuring multi-bot pool rotation, AIMD upload pacing, and in-place HatchZip extraction.</i>
</p>

</div>

---

> [!IMPORTANT]
> **Pure Leech Bot Notice**: `AZMLX-zone` is exclusively designed for **Leeching** (downloading files from direct links, torrents, yt-dlp, and Telegram to upload straight back to Telegram). **Cloud Drive Mirroring (Google Drive, RClone, Mega) is not included for now** to maintain absolute focus on raw leeching speed, stability, and minimal resource usage.

---
| Command | Description |
|---|---|
| `/leech [link/reply]` | Download and leech to Telegram |
| `/unzipleech [link/reply]` | Download, extract archive in-place, and leech contents |
| `/zipleech [link/reply]` | Download, compress into `.zip`, and leech |
| `/ytdl [link]` | Download video/audio via yt-dlp and leech to Telegram |
| `/ytdlzip [link]` | Download via yt-dlp, package as zip, and leech |
| `/status` | View real-time progress, speed, ETA, and queue status |
| `/cancel [gid]` | Cancel a specific running task |
| `/cancelall` | Cancel all running tasks |
| `/settings` | Open interactive user settings menu (Thumbnails, Prefixes, Fonts) |
| `/speedtest` | Test server download and upload bandwidth |
| `/ping` | Check bot responsiveness and latency |
| `/stats` | View bot and system resource usage |
| `/help` | Display command usage and instructions |

---

## 🛠️ Deployment Methods

Choose between **Heroku CLI (Server-side Build)**, **Automated GitHub Actions CD**, or **Linux VPS (Docker Compose)**.

---

### Method 1: Deploy on Heroku (Heroku CLI - Server Build)

Deploys `AZML-Go` using Heroku's **Container Stack** where Heroku builds the `Dockerfile` directly on their remote build servers via `heroku.yml` (no local Docker build required).

#### 1. Prerequisites
- Install the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli):
  ```bash
  curl https://cli-assets.heroku.com/install.sh | sh
  ```

#### 2. Clone the Repository
```bash
git clone https://github.com/aquib4040/azml-go.git
cd azml-go
```

#### 3. Log In to Heroku
```bash
heroku login
```

#### 4. Create Heroku App with Container Stack
```bash
# If creating a new app:
heroku apps:create your-app-name --stack=container

# Or if linking an existing app:
heroku git:remote -a your-app-name
heroku stack:set container -a your-app-name
```

#### 5. Configure Environment Variables
Set your configuration variables on Heroku:
```bash
# Grab your live App Web URL to use as BASE_URL
APP_URL=$(heroku apps:info -a your-app-name --json | jq -r '.app.web_url')

# Set configuration variables
heroku config:set BOT_TOKEN="your_bot_token" -a your-app-name
heroku config:set OWNER_ID="your_telegram_id" -a your-app-name
heroku config:set TELEGRAM_API="your_api_id" -a your-app-name
heroku config:set TELEGRAM_HASH="your_api_hash" -a your-app-name
heroku config:set BASE_URL="$APP_URL" -a your-app-name
```

#### 6. Deploy Code (Remote Docker Build on Heroku)
When pushing to Heroku git, Heroku detects `heroku.yml` and builds the Docker container directly on their servers:
```bash
git push heroku master
```

#### 7. View Live Logs
```bash
heroku logs --tail -a your-app-name
```

#### 8. Authors

Thanks to the code owner @aquib4040


