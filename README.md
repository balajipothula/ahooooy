# 🚀 GoFiber App — One-Click Deployment into Fly.io 

✨ A dynamic, feathery [GoFiber](https://gofiber.io/) web application, ready to **fly** with just one click!  
Thanks to [GitHub Actions](.github/workflows/cd_fly-io_app_deployment.yml), this repo enables **fully automated deployment** to [Fly.io](https://fly.io) — no manual steps required.

---

## 🌟 Features

- **One-Click Deployment** — Simply push to `main` and GitHub Actions handles the rest.  
- **Idempotent Pipeline** — Safe to re-run without double-creating volumes.  
- **Self-Healing** — If a volume is deleted, the next deploy automatically recreates it.  
- **No Manual CLI Steps** — Everything runs directly inside your GitHub workflow.
- **Database Change Management & CI/CD** — Tracking, managing and applying database schema changes using Liquibase GitHub Action.  
- **Fast & Minimalistic** — API powered by GoFiber.  
- **Scalable Design** — Modular folder structure for growth.  
- **Database Ready** — PostgreSQL integration via GORM.
- **Best Practices** — Built for reproducibility and automation.  

---

## ⚡️ Quick Start

1. **Fork this repo** or clone it:
   ```bash
   git clone https://github.com/balajipothula/go-fiber-app.git
   cd go-fiber-app
### 2. Configure Fly.io
- Create a [Fly.io](https://fly.io) account  
- Generate an API token  
- Add the token as a GitHub Secret named **`FLY_API_TOKEN`**

### 3. Push to Deploy 🚀
- Commit & push changes to the `main` branch  
- GitHub Actions will automatically:
  - Build & push the Docker image  
  - Manage Fly.io volumes  
  - Deploy the app to Fly.io  
