<template>
  <div class="flex h-screen w-screen overflow-hidden bg-gray-100">
    <!-- Sidebar -->
    <aside 
      :class="[isCollapsed ? 'w-16' : 'w-64']"
      class="bg-gray-900 text-white transition-all duration-300 flex flex-col"
    >
      <!-- Header / Toggle -->
      <div class="p-4 flex items-center justify-between border-b border-gray-800">
        <span v-show="!isCollapsed" class="font-bold text-lg">App Title</span>
        <button @click="isCollapsed = !isCollapsed" class="p-2 rounded hover:bg-gray-800">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
        </button>
      </div>

      <!-- Navigation Links -->
      <nav class="flex-1 p-4 space-y-2">
        <router-link 
          to="/" 
          class="flex items-center p-2 rounded hover:bg-gray-800"
          :class="{'bg-blue-600': $route.path === '/'}"
        >
          <svg class="w-5 h-5 min-w-[20px]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          <span v-show="!isCollapsed" class="ml-4 transition-opacity">Home</span>
        </router-link>
        
        <router-link 
          to="/settings" 
          class="flex items-center p-2 rounded hover:bg-gray-800"
          :class="{'bg-blue-600': $route.path === '/settings'}"
        >
          <svg class="w-5 h-5 min-w-[20px]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span v-show="!isCollapsed" class="ml-4 transition-opacity">Settings</span>
        </router-link>
      </nav>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 overflow-y-auto p-6 bg-white">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isCollapsed = ref(false)
</script>
