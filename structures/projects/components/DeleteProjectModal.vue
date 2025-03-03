<template>
  <UModal v-model:open="isOpen">
    <UCard>
      <template #header>
        <div class="flex items-center justify-between">
          <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
            Delete Project
          </h3>
          <UButton
            color="neutral"
            variant="ghost"
            icon="i-heroicons-x-mark-20-solid"
            class="mr-1"
            @click="isOpen = false"
          />
        </div>
      </template>
      
      <template #body>
        <div class="p-4">
          <p class="text-sm text-gray-500 dark:text-gray-400">
            You're deleting the following Project. This action cannot be undone.
          </p>
          <p v-if="item" class="mt-2 font-medium">
            {{ item.name }}
          </p>
        </div>
      </template>
      
      <template #footer>
        <div class="flex justify-end gap-3">
          <UButton
            color="neutral"
            variant="soft"
            @click="isOpen = false"
          >
            Cancel
          </UButton>
          <UButton
            color="error"
            :loading="loading"
            @click="confirmDelete"
          >
            Confirm
          </UButton>
        </div>
      </template>
    </UCard>
  </UModal>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import type { Project } from '../stores/project'

const props = defineProps<{
  open?: boolean
  project?: Project
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

const item = computed(() => props.project)

function confirmDelete() {
  emit('confirm')
}
</script>
