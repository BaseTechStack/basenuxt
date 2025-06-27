<template>
  <UModal v-model:open="isOpen" title="Are you sure?" description="You are about to delete this winetype." :ui="{ footer: 'justify-end ', content: 'text-white dark:text-white bg-red-500 dark:bg-red-900', }">
 
    <template #body>
      <div class="space-y-1">
        <div class="text-sm font-medium text-white dark:text-white">Name</div>
        <div>{{ winetype?.name }}</div>
      </div>
 
    </template>

    <template #footer>
      <UButton label="Cancel" color="error" variant="outline" @click="isOpen = false" />
      <UButton label="Delete" color="error" :loading="loading" @click="confirmDelete" />
    </template>
  </UModal>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import type { WineType } from '../stores/winetype'

const props = defineProps<{
  open?: boolean
  winetype?: WineType
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'confirm'): void
}>()

const isOpen = computed({
  get() {
    return props.open
  },
  set(value) {
    emit('update:open', value)
  }
})

const loading = computed(() => props.loading || false)



function confirmDelete() {
  emit('confirm')
}
</script>
