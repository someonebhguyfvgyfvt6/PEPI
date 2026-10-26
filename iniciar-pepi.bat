@echo off
cd /d "%~dp0pepi-frontend"
if not exist node_modules call npm install
call npm run dev -- --open
