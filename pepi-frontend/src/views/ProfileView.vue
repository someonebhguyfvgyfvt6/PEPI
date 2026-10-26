<script setup lang="ts">
import ProjectCard from '@/components/ProjectCard.vue'
</script>

<template>
  <div class="mx-auto max-w-6xl px-4 py-8">
    <div class="card overflow-hidden">
      <div class="ph h-40 rounded-none"></div>
      <div class="flex flex-wrap items-end justify-between gap-4 px-6 pb-6">
        <div class="flex items-end gap-4">
          <div class="-mt-10 h-24 w-24 rounded-full border-4 border-white bg-neutral-300"></div>
          <div>
            <div class="text-xl font-bold">Nombre del freelancer</div>
            <div class="text-sm text-neutral-500">Oficio · País</div>
          </div>
        </div>
        <div class="flex gap-2">
          <RouterLink to="/mensajes" class="btn">Mensaje</RouterLink>
          <button class="btn btn-primary">Contratar</button>
          <button class="btn">Descargar CV</button>
        </div>
      </div>
    </div>

    <div class="mt-6 grid gap-6 md:grid-cols-[1fr_280px]">
      <section>
        <h2 class="mb-3 font-semibold">Proyectos</h2>
        <div class="grid grid-cols-2 gap-4 lg:grid-cols-3">
          <ProjectCard v-for="n in 6" :key="n" />
        </div>
      </section>
      <aside class="space-y-4">
        <div class="card p-4">
          <h3 class="mb-2 text-sm font-semibold">Sobre mí</h3>
          <div class="space-y-2"><div class="h-3 rounded bg-neutral-200"></div><div class="h-3 w-4/5 rounded bg-neutral-200"></div><div class="h-3 w-2/3 rounded bg-neutral-200"></div></div>
        </div>
        <div class="card p-4">
          <h3 class="mb-2 text-sm font-semibold">Especialidades</h3>
          <div class="flex flex-wrap gap-2"><span v-for="n in 4" :key="n" class="h-6 w-16 rounded-full bg-neutral-200"></span></div>
        </div>
        <div class="card p-4">
          <h3 class="mb-2 text-sm font-semibold">Reseñas</h3>
          <div class="space-y-2"><div class="h-3 rounded bg-neutral-200"></div><div class="h-3 w-3/4 rounded bg-neutral-200"></div></div>
        </div>
      </aside>
    </div>
  </div>
</template>
