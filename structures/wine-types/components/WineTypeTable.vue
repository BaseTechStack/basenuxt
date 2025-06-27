<template>
  <div class="p-4">
    <USkeleton v-if="loading" class="h-96" />
    
    <UAlert v-else-if="error" :description="error" color="error" variant="soft" />
    
    <div v-else>
      <div class="flex items-center gap-2 p-2 mb-4">
        <UInput
          v-model="search"
          icon="i-heroicons-magnifying-glass"
          placeholder="Search winetypes..."
          size="sm"
          class="max-w-sm"
        />
      </div>
      
      <UTable :data="winetypes" :columns="columns" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { WineType } from '../stores/winetype'
import { ref, h, resolveComponent, computed } from 'vue'
import type { TableColumn } from '#imports'

interface Props {
  winetype: WineType[]
  loading?: boolean
  error?: string | null
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  error: null
})

const emit = defineEmits<{
  (e: 'edit', winetype: WineType): void
  (e: 'delete', winetype: WineType): void
  (e: 'view', winetype: WineType): void
}>()

const search = ref('')

const winetypes = computed(() => props.winetype)

const columns: TableColumn<WineType>[] = [
  {
    accessorKey: 'id',
    header: 'ID'
  },
  
  {
    accessorKey: 'name',
    header: 'Name'
  },
  
  {
    accessorKey: 'description',
    header: 'Description'
  },
  
  {
    accessorKey: 'category',
    header: 'Category'
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
        onClick: () => emit('view', row as unknown as WineType)
      }),
      h(resolveComponent('UButton'), {
        color: 'primary',
        variant: 'ghost',
        icon: 'i-heroicons-pencil-square',
        size: 'xs',
        onClick: () => emit('edit', row as unknown as WineType)
      }),
      h(resolveComponent('UButton'), {
        color: 'red',
        variant: 'ghost',
        icon: 'i-heroicons-trash',
        size: 'xs',
        onClick: () => emit('delete', row as unknown as WineType)
      })
    ])
  }
]
</script>
