<script setup lang="ts">
import PageTitle from '@/components/PageTitle.vue'
import { useUiStore } from '@/stores/ui'
const ui = useUiStore()
const stats = ['Visitas al perfil', 'Mensajes nuevos', 'Proyectos activos']
</script>

<template>
  <div class="mx-auto max-w-6xl px-4 py-8">
    <PageTitle :title="ui.rol === 'empresa' ? 'Panel de empresa' : 'Panel de freelancer'" />
    <div class="mb-6 inline-flex rounded-lg border border-neutral-300 bg-white p-1 text-sm">
      <button v-for="r in ['freelancer','empresa'] as const" :key="r" class="rounded-md px-3 py-1"
        :class="ui.rol === r ? 'bg-neutral-900 text-white' : 'text-neutral-600'" @click="ui.setRol(r)">
        {{ r === 'empresa' ? 'Empresa' : 'Freelancer' }}
      </button>
    </div>

    <div class="grid gap-4 md:grid-cols-3">
      <div v-for="s in stats" :key="s" class="card p-4">
        <div class="text-sm text-neutral-500">{{ s }}</div>
        <div class="mt-2 h-7 w-16 rounded bg-neutral-200"></div>
      </div>
    </div>

    <div class="mt-6 grid gap-6 md:grid-cols-2">
      <section class="card p-4">
        <h2 class="mb-3 font-semibold">{{ ui.rol === 'empresa' ? 'Vacantes' : 'Mis proyectos' }}</h2>
        <div class="space-y-3"><div v-for="n in 4" :key="n" class="ph h-12"></div></div>
        <button class="btn mt-4">{{ ui.rol === 'empresa' ? 'Publicar vacante' : 'Añadir proyecto' }}</button>
      </section>
      <section class="card p-4">
        <h2 class="mb-3 font-semibold">{{ ui.rol === 'empresa' ? 'Talento guardado' : 'Propuestas' }}</h2>
        <div class="space-y-3"><div v-for="n in 4" :key="n" class="ph h-12"></div></div>
      </section>
    </div>
  </div>
</template>
