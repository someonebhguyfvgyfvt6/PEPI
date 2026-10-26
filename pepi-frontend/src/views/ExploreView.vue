<script setup lang="ts">
import PageTitle from '@/components/PageTitle.vue'
import ProjectCard from '@/components/ProjectCard.vue'
import { useUiStore } from '@/stores/ui'

const ui = useUiStore()
const categorias = [
  'Tecnología y programación', 'Diseño', 'Marketing', 'Escritura e idiomas',
  'Foto, video y audio', 'Administración', 'Ventas y e-commerce',
  'Arquitectura e ingeniería', 'Educación', 'Música', 'Servicios especializados',
]
</script>

<template>
  <div class="mx-auto max-w-6xl px-4 py-8">
    <PageTitle title="Explorar" subtitle="Busca talento por categoría y especialidad." />
    <input class="input mb-6" placeholder="Buscar por nombre, oficio o habilidad" />
    <div class="grid gap-6 md:grid-cols-[240px_1fr]">
      <aside class="card h-fit p-4">
        <h2 class="mb-3 text-sm font-semibold">Categorías</h2>
        <ul class="space-y-1">
          <li v-for="c in categorias" :key="c">
            <button class="w-full rounded-md px-2 py-1.5 text-left text-sm hover:bg-neutral-100"
              :class="ui.categoria === c ? 'bg-neutral-100 font-semibold' : 'text-neutral-600'"
              @click="ui.setCategoria(ui.categoria === c ? null : c)">{{ c }}</button>
          </li>
        </ul>
        <h2 class="mb-3 mt-5 text-sm font-semibold">Disponibilidad</h2>
        <label class="flex items-center gap-2 text-sm text-neutral-600"><input type="checkbox" /> Disponible ahora</label>
      </aside>
      <div class="grid grid-cols-2 gap-4 lg:grid-cols-3">
        <ProjectCard v-for="n in 9" :key="n" />
      </div>
    </div>
  </div>
</template>
