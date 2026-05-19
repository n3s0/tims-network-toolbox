<template>
  <aside :class="['sidebar', { collapsed: isCollapsed }]">
    <button class="toggle-btn" @click="toggleSidebar">
      {{ isCollapsed ? '▶' : '◀' }}
    </button>
    <nav class="nav-links">
      <router-link to="/dashboard">Dashboard</router-link>
      <router-link to="/settings">Settings</router-link>
    </nav>
  </aside>
</template>

<script setup>
import { ref } from 'vue';

const isCollapsed = ref(false);

const toggleSidebar = () => {
  isCollapsed.value = !isCollapsed.value;
};
</script>

<style scoped>
.sidebar {
  width: 250px;
  height: 100vh;
  background-color: #2c3e50;
  transition: width 0.3s ease;
  display: flex;
  flex-direction: column;
}

.sidebar.collapsed {
  width: 60px;
}

.toggle-btn {
  align-self: flex-end;
  margin: 10px;
  cursor: pointer;
}

.nav-links {
  display: flex;
  flex-direction: column;
  padding: 20px 0;
}

.nav-links a {
  color: white;
  text-decoration: none;
  padding: 15px;
  white-space: nowrap;
}

.sidebar.collapsed .nav-links a {
  /* Hide text or just show icons when collapsed */
}
</style>
