import { defineStore } from 'pinia'

export type Rol = 'freelancer' | 'empresa'

export const useUiStore = defineStore('ui', {
  state: () => ({
    rol: 'freelancer' as Rol,
    categoria: null as string | null,
  }),
  actions: {
    setRol(r: Rol) { this.rol = r },
    setCategoria(c: string | null) { this.categoria = c },
  },
})
