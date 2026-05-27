# Screenshots & Proof Artifacts

This folder holds the visual proof that the platform actually runs. Drop the
images in here with the exact filenames below and they will render automatically
in the root `README.md`.

> These are placeholders. Replace each one with a real screenshot from your own
> running environment.

## Required screenshots

| Filename | What to capture | How to get there |
|----------|-----------------|------------------|
| `grafana-dashboard.png` | The **Microservices** Grafana dashboard showing request rate, latency (p95), and error rate panels with live data. | `make dev-up`, open http://localhost:3000 (login `admin` / `admin`), open the *Microservices* dashboard, generate some traffic with a few `curl` calls, then screenshot. |
| `argocd-sync.png` | The **ArgoCD** application view showing the apps **Synced** and **Healthy** (the app-of-apps tree is ideal). | Port-forward the ArgoCD server (`kubectl port-forward svc/argocd-server -n argocd 8080:443`), open the UI, and screenshot the applications list or a single app's resource tree after a sync. |
| `github-actions-run.png` | A **passing** GitHub Actions run showing the `test`, `lint`, `security-scan`, and `build-and-push` jobs green. | Push to `main`, open the Actions tab, open the completed run, and screenshot the job graph. |

## Tips for good screenshots

- Use a clean browser window (no unrelated tabs/bookmarks bar) or crop to the panel.
- Make sure there is **real data** on the Grafana panels — hit the API a few times
  first so the graphs aren't empty.
- For ArgoCD, the green **Synced / Healthy** badges are the money shot.
- Keep images reasonably sized (PNG, ~1600px wide is plenty) so the README stays light.

## Optional extras

- `prometheus-targets.png` — Prometheus *Targets* page showing all services `UP`.
- `terraform-apply.png` — a successful `terraform apply` summary.
