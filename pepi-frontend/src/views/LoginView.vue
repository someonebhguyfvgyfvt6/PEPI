<template>
  <div class="mx-auto max-w-sm px-4 py-16">
    <div class="card p-6">
      <h1 class="mb-6 text-xl font-bold">Entrar a PEPI</h1>
      <form class="space-y-4" @submit.prevent>
        <div><label class="mb-1 block text-sm">Correo</label><input class="input" type="email" /></div>
        <div><label class="mb-1 block text-sm">Contraseña</label><input class="input" type="password" /></div>
        <button class="btn btn-primary w-full">Entrar</button>
      </form>
      <p class="mt-4 text-center text-sm text-neutral-500">
        ¿No tienes cuenta? <RouterLink to="/registro" class="font-medium text-neutral-900 underline">Crear cuenta</RouterLink>
      </p>
    </div>
  </div>
</template>
