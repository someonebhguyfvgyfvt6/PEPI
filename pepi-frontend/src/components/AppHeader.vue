<script setup lang="ts">
const links = [
  { to: '/explorar', label: 'Explorar' },
  { to: '/mensajes', label: 'Mensajes' },
  { to: '/panel', label: 'Panel' },
  { to: '/perfil', label: 'Mi perfil' },
]
</script>

<template>
  <header class="border-b border-neutral-200 bg-white">
    <div class="mx-auto flex h-14 max-w-6xl items-center justify-between px-4">
      <RouterLink to="/" class="text-lg font-bold tracking-tight">PEPI</RouterLink>
      <nav class="hidden items-center gap-6 text-sm text-neutral-600 sm:flex">
        <RouterLink v-for="l in links" :key="l.to" :to="l.to" class="hover:text-neutral-900"
          active-class="font-semibold text-neutral-900">{{ l.label }}</RouterLink>
      </nav>
      <div class="flex items-center gap-2">
        <RouterLink to="/entrar" class="btn">Entrar</RouterLink>
        <RouterLink to="/registro" class="btn btn-primary">Crear cuenta</RouterLink>
      </div>
    </div>
  </header>
</template>
