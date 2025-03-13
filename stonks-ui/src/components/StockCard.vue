<template>
  <div class="card">
    <div 
      class="card-header"
      @click="toggleExpand"
    >
      <div>
        <div class="flex items-center">
          <span class="text-2xl font-bold text-blue-700 mr-3">{{ stock.ticker }}</span>
          <span class="text-gray-600 text-sm bg-blue-100 px-2 py-1 rounded">
            {{ stock.company }} - {{ stock.brokerage }}
          </span>
        </div>
        <p class="text-sm text-gray-500 mt-2">
          {{ formatDate(stock.time) }}
        </p>
      </div>
      <div>
        <span class="transform transition-transform duration-300" :class="{ 'rotate-180': isExpanded }">
          ▼
        </span>
      </div>
    </div>

    <div 
      v-if="isExpanded" 
      class="px-6 py-4 bg-gray-50 border-t border-gray-200"
    >
      <div class="grid grid-cols-2 gap-4">
        <div>
          <span class="text-sm font-medium text-gray-600">Brokerage</span>
          <p class="text-gray-800 font-semibold">{{ stock.brokerage }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Action</span>
          <p class="text-gray-800 font-semibold">{{ stock.action }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Rating From</span>
          <p class="text-gray-800 font-semibold">{{ stock.rating_from }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Rating To</span>
          <p class="text-gray-800 font-semibold">{{ stock.rating_to }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Target From</span>
          <p class="text-green-600 font-bold">${{ stock.target_from.toFixed(2) }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Target To</span>
          <p class="text-green-600 font-bold">${{ stock.target_to.toFixed(2) }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Stock } from '@/services/api';

// eslint-disable-next-line no-unused-vars
const props = defineProps<{
  stock: Stock;
}>();

const isExpanded = ref(false);

const toggleExpand = () => {
  isExpanded.value = !isExpanded.value;
};

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
};
</script>