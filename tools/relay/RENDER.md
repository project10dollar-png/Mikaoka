# Mika Relay on Render

This repository contains a Render-ready copy of the optional Mika Relay service from the upstream Mika project.

## Deployment

The root `render.yaml` defines one Docker web service in the Frankfurt region. It uses a 1 GB persistent disk at `/data` for relay pairing state and a Render-generated `MIKA_OWNER_KEY` secret. The relay reads the public HTTPS origin from `RENDER_EXTERNAL_URL` and binds to Render's `PORT`.

The persistent disk requires a paid Render web-service plan and disables zero-downtime deploys. Review the cost and resource plan in Render before approving the Blueprint.

## Security

- Keep `MIKA_OWNER_KEY` secret. Do not paste it into chat, commit it, or expose it in logs.
- The gateway is a remote access endpoint. Only pair devices you control.
- TLS is terminated by Render's HTTPS edge.
- The gateway forwards requests and therefore can process their content in memory; it is not end-to-end encrypted.
- The original upstream repository is not modified by this deployment copy.

## Upstream

Source: https://github.com/wangzan101/Mika/tree/main/tools/relay
