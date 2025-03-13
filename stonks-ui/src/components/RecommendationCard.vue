<template>
  <div class="card">
    <div 
      class="card-header"
      @click="toggleExpand"
    >
      <div>
        <div class="flex items-center">
          <span class="text-2xl font-bold text-blue-700 mr-3">{{ recommendation.stock.ticker }}</span>
          <span class="text-gray-600 text-sm bg-blue-100 px-2 py-1 rounded">
            {{ recommendation.stock.company }} - {{ recommendation.stock.brokerage }}
          </span>
        </div>
        <div class="flex items-center mt-2">
          <div class="px-3 py-1 rounded text-sm font-medium"
            :class="getScoreClass(recommendation.score)">
            {{ getScoreLabel(recommendation.score) }}
          </div>
        </div>
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
      <div class="mb-4">
        <span class="text-sm font-medium text-gray-600">Recommendation Reason</span>
        <p class="text-gray-800 font-semibold">{{ recommendation.reason }}</p>
      </div>
      
      <div class="grid grid-cols-2 gap-4">
        <div>
          <span class="text-sm font-medium text-gray-600">Brokerage</span>
          <p class="text-gray-800 font-semibold">{{ recommendation.stock.brokerage }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Action</span>
          <p class="text-gray-800 font-semibold">{{ recommendation.stock.action }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Rating From</span>
          <p class="text-gray-800 font-semibold">{{ recommendation.stock.rating_from }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Rating To</span>
          <p class="text-gray-800 font-semibold">{{ recommendation.stock.rating_to }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Target From</span>
          <p class="text-green-600 font-bold">${{ recommendation.stock.target_from.toFixed(2) }}</p>
        </div>
        <div>
          <span class="text-sm font-medium text-gray-600">Target To</span>
          <p class="text-green-600 font-bold">${{ recommendation.stock.target_to.toFixed(2) }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { StockRecommendation } from '@/services/api';

// eslint-disable-next-line no-unused-vars
const props = defineProps<{
  recommendation: StockRecommendation;
}>();

const isExpanded = ref(false);

const toggleExpand = () => {
  isExpanded.value = !isExpanded.value;
};

const getScoreClass = (score: number) => {
  if (score >= 4) return 'bg-green-100 text-green-700';
  if (score >= 2.5) return 'bg-blue-100 text-blue-700';
  if (score >= 1) return 'bg-yellow-100 text-yellow-700';
  return 'bg-red-100 text-red-700';
};

const getScoreLabel = (score: number) => {
  if (score >= 4) return 'Strong Buy';
  if (score >= 2.5) return 'Buy';
  if (score >= 1) return 'Hold';
  return 'Sell';
};
</script>