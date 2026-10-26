<script setup lang="ts">
import ProjectCard from '@/components/ProjectCard.vue'
</script>

<template>
  <section class="mx-auto max-w-6xl px-4 py-16 text-center">
    <h1 class="text-4xl font-bold tracking-tight sm:text-5xl">Tu trabajo, tu casa.</h1>
    <p class="mx-auto mt-4 max-w-xl text-neutral-600">
      PEPI es la red donde tu portafolio es tu perfil. Muestra lo que haces y deja que te encuentren por tu trabajo.
    </p>
    <div class="mt-8 flex justify-center gap-3">
      <RouterLink to="/registro" class="btn btn-primary">Crear mi portafolio</RouterLink>
      <RouterLink to="/explorar" class="btn">Explorar talento</RouterLink>
    </div>
  </section>

  <section class="mx-auto max-w-6xl px-4 pb-16">
    <div class="grid grid-cols-2 gap-4 md:grid-cols-4">
      <ProjectCard v-for="n in 8" :key="n" />
    </div>
  </section>

  <section class="border-t border-neutral-200 bg-white">
    <div class="mx-auto grid max-w-6xl gap-6 px-4 py-12 md:grid-cols-3">
      <div v-for="t in ['Para freelancers','Para empresas','Para equipos']" :key="t" class="card p-6">
        <div class="ph mb-4 h-10 w-10"></div>
        <h3 class="font-semibold">{{ t }}</h3>
        <div class="mt-3 space-y-2">
          <div class="h-3 rounded bg-neutral-200"></div>
          <div class="h-3 w-5/6 rounded bg-neutral-200"></div>
        </div>
      </div>
    </div>
  </section>
</template>
