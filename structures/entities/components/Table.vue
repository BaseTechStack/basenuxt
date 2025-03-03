<template>
  <div class="p-4">
    <USkeleton v-if="loading" class="h-96" />
    
    <UAlert v-else-if="error" :description="error" color="error" variant="soft" />
    
    <div v-else>
      <div class="flex items-center gap-2 p-2 mb-4">
        <UInput
          v-model="search"
          icon="i-heroicons-magnifying-glass"
          placeholder="Search entities..."
          size="sm"
          class="max-w-sm"
        />
      </div>
      
      <UTable :data="entities" :columns="columns" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { entity } from '../stores/entity'
import { ref, h, resolveComponent, computed } from 'vue'
import type { TableColumn } from '@nuxt/ui'

interface Props {
  entity: entity[]
  loading?: boolean
  error?: string | null
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  error: null
})

const emit = defineEmits<{
  (e: 'edit', entity: entity): void
  (e: 'delete', entity: entity): void
  (e: 'view', entity: entity): void
}>()

const search = ref('')

const entities = computed(() => props.entity)

const columns: TableColumn<entity>[] = [
  {
    accessorKey: 'id',
    header: 'ID'
  },
  
  {
    accessorKey: 'test',
    header: 'Test'
  },
  
  {
    accessorKey: 'name',
    header: 'Name'
  },
  
  {
    accessorKey: 'email',
    header: 'Email'
  },
  
  {
    accessorKey: 'age',
    header: 'Age'
  },
  
  {
    accessorKey: 'active',
    header: 'Active'
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
        onClick: () => emit('view', row as unknown as entity)
      }),
      h(resolveComponent('UButton'), {
        color: 'primary',
        variant: 'ghost',
        icon: 'i-heroicons-pencil-square',
        size: 'xs',
        onClick: () => emit('edit', row as unknown as entity)
      }),
      h(resolveComponent('UButton'), {
        color: 'red',
        variant: 'ghost',
        icon: 'i-heroicons-trash',
        size: 'xs',
        onClick: () => emit('delete', row as unknown as entity)
      })
    ])
  }
]
</script>
