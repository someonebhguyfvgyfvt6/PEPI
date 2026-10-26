import axios from 'axios'

// Cliente base para el backend en Go (Fiber). En la Fase 0 aún no se usa.
export const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL ?? 'http://localhost:3000/api',
})
