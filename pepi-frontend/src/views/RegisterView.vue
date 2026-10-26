<script setup lang="ts">
import { useUiStore } from '@/stores/ui'
const ui = useUiStore()
const roles = [
  { id: 'freelancer', t: 'Soy freelancer', d: 'Quiero mostrar mi portafolio.' },
  { id: 'empresa', t: 'Soy empresa', d: 'Quiero encontrar talento.' },
] as const
</script>

<template>
  <div class="mx-auto max-w-md px-4 py-16">
    <div class="card p-6">
      <h1 class="mb-1 text-xl font-bold">Crear cuenta</h1>
      <p class="mb-5 text-sm text-neutral-500">Elige cómo vas a usar PEPI.</p>
      <div class="mb-5 grid grid-cols-2 gap-3">
        <button v-for="r in roles" :key="r.id" type="button" @click="ui.setRol(r.id)"
          class="rounded-xl border p-4 text-left"
          :class="ui.rol === r.id ? 'border-neutral-900 bg-neutral-100' : 'border-neutral-300 bg-white'">
          <div class="text-sm font-semibold">{{ r.t }}</div>
          <div class="mt-1 text-xs text-neutral-500">{{ r.d }}</div>
        </button>
      </div>
      <form class="space-y-4" @submit.prevent>
        <div><label class="mb-1 block text-sm">{{ ui.rol === 'empresa' ? 'Nombre de la empresa' : 'Nombre' }}</label><input class="input" /></div>
        <div><label class="mb-1 block text-sm">Correo</label><input class="input" type="email" /></div>
        <div><label class="mb-1 block text-sm">Contraseña</label><input class="input" type="password" /></div>
        <button class="btn btn-primary w-full">Crear cuenta</button>
      </form>
      <p class="mt-4 text-center text-sm text-neutral-500">
        ¿Ya tienes cuenta? <RouterLink to="/entrar" class="font-medium text-neutral-900 underline">Entrar</RouterLink>
      </p>
    </div>
  </div>
</template>
