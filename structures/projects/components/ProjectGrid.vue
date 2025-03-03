<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
    <ProjectGridCard 
      v-for="project in projects" 
      :key="project.id" 
      :item="project"
      title="id"
      subtitle="/users/flakerimismani/base/basenuxt/testProject"
      :fields="[
        
        
        
        
        
      ]"
      timestamp="createdAt"
      hover
    >
      <template #actions>
        <BaseCrudActions 
          structure="projects"
          :item="project" 
          actions="view,edit,delete"
          @view="$emit('view', $event)"
          @edit="$emit('edit', $event)"
          @delete="$emit('delete', $event)"
        />
      </template>
      
      <template #footer-actions>
        <UButton
          size="xs"
          color="primary"
          variant="ghost"
          icon="i-heroicons-eye"
          @click="$emit('view', project)"
        >
          View Details
        </UButton>
      </template>
    </ProjectGridCard>
  </div>
</template>

<script setup lang="ts">
import type { Project } from '../stores/project'
import { computed } from 'vue'

const props = defineProps<{
  project: Project[]
}>()

const projects = computed(() => props.project)

defineEmits<{
  (e: 'view', project: Project): void
  (e: 'edit', project: Project): void
  (e: 'delete', project: Project): void
}>()
</script>
