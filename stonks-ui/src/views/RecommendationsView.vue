<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <div v-if="loading" class="text-center py-8">
      <div class="animate-pulse text-gray-500">Loading recommendations...</div>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4" role="alert">
      {{ error }}
    </div>

    <div v-else-if="recommendations.length === 0" class="text-center py-8 bg-gray-100 rounded-lg shadow-inner">
      <p class="text-gray-700 font-medium">No recommendations available at this time</p>
      <p class="text-gray-500 text-sm mt-2">Check back later for updated investment suggestions</p>
    </div>

    <div v-else class="space-y-4">
      <div class="bg-blue-50 border-l-4 border-blue-500 p-4 mb-4">
        <p class="text-blue-700">
          We found <span class="font-bold">{{ recommendations.length }}</span> recommended stocks for you
        </p>
      </div>
      
      <RecommendationCard 
        v-for="rec in recommendations" 
        :key="rec.stock.id" 
        :recommendation="rec" 
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { storeToRefs } from 'pinia';
import { useRecommendationsStore } from '@/stores/recommendations';
import RecommendationCard from '@/components/RecommendationCard.vue';

// Use storeToRefs to extract reactive properties from the store
const recommendationsStore = useRecommendationsStore();
const { recommendations, loading, error } = storeToRefs(recommendationsStore);

onMounted(() => {
  recommendationsStore.fetchRecommendations();
});
</script>