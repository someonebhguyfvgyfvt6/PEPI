<script setup lang="ts">
import PageTitle from '@/components/PageTitle.vue'
</script>

<template>
  <div class="mx-auto max-w-6xl px-4 py-8">
    <PageTitle title="Mensajes" />
    <div class="card grid h-[520px] overflow-hidden md:grid-cols-[280px_1fr]">
      <aside class="border-r border-neutral-200">
        <div class="border-b border-neutral-200 p-3"><input class="input" placeholder="Buscar conversación" /></div>
        <ul>
          <li v-for="n in 6" :key="n" class="flex items-center gap-3 border-b border-neutral-100 p-3" :class="n === 1 && 'bg-neutral-100'">
            <div class="h-10 w-10 shrink-0 rounded-full bg-neutral-300"></div>
            <div class="flex-1 space-y-2"><div class="h-3 w-1/2 rounded bg-neutral-200"></div><div class="h-3 w-4/5 rounded bg-neutral-200"></div></div>
          </li>
        </ul>
      </aside>
      <section class="flex flex-col">
        <div class="flex items-center gap-3 border-b border-neutral-200 p-3">
          <div class="h-8 w-8 rounded-full bg-neutral-300"></div><div class="h-3 w-32 rounded bg-neutral-200"></div>
        </div>
        <div class="flex-1 space-y-3 p-4">
          <div class="h-8 w-1/2 rounded-lg bg-neutral-200"></div>
          <div class="ml-auto h-8 w-2/5 rounded-lg bg-neutral-800"></div>
          <div class="h-8 w-1/3 rounded-lg bg-neutral-200"></div>
        </div>
        <form class="flex gap-2 border-t border-neutral-200 p-3" @submit.prevent>
          <input class="input" placeholder="Escribe un mensaje" /><button class="btn btn-primary">Enviar</button>
        </form>
      </section>
    </div>
  </div>
</template>
