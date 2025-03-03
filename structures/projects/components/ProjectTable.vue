<template>
  <div class="p-4">
    <USkeleton v-if="loading" class="h-96" />
    
    <UAlert v-else-if="error" :description="error" color="error" variant="soft" />
    
    <div v-else>
      <div class="flex items-center gap-2 p-2 mb-4">
        <UInput
          v-model="search"
          icon="i-heroicons-magnifying-glass"
          placeholder="Search projects..."
          size="sm"
          class="max-w-sm"
        />
      </div>
      
      <UTable :data="projects" :columns="columns" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Project } from '../stores/project'
import { ref, h, resolveComponent, computed } from 'vue'
import type { TableColumn } from '@nuxt/ui'

interface Props {
  project: Project[]
  loading?: boolean
  error?: string | null
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  error: null
})

const emit = defineEmits<{
  (e: 'edit', project: Project): void
  (e: 'delete', project: Project): void
  (e: 'view', project: Project): void
}>()

const search = ref('')

const projects = computed(() => props.project)

const columns: TableColumn<Project>[] = [
  {
    accessorKey: 'id',
    header: 'ID'
  },
  
  {
    accessorKey: 'name',
    header: 'Name'
  },
  
  {
    accessorKey: '/users/flakerimismani/base/basenuxt/testProject',
    header: '/users/flakerimismani/base/basenuxt/testproject'
  },
  
  {
    accessorKey: 'createdAt',
    header: 'Created'
  },
  {
    accessorKey: 'actions',
    header: () => h('div', { class: 'text-right' }, 'Actions'),
    cell: ({ row }) => h('div', { class: 'flex justify-end gap-2' }, [
      h(resolveComponent('UButton'), {
        color: 'primary',
        variant: 'ghost',
        icon: 'i-heroicons-eye',
        size: 'xs',
        onClick: () => emit('view', row as unknown as Project)
      }),
      h(resolveComponent('UButton'), {
        color: 'primary',
        variant: 'ghost',
        icon: 'i-heroicons-pencil-square',
        size: 'xs',
        onClick: () => emit('edit', row as unknown as Project)
      }),
      h(resolveComponent('UButton'), {
        color: 'red',
        variant: 'ghost',
        icon: 'i-heroicons-trash',
        size: 'xs',
        onClick: () => emit('delete', row as unknown as Project)
      })
    ])
  }
]
</script>
