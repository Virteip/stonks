<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <div class="mb-8 flex items-center justify-between">
      <div class="relative flex-grow mr-4">
        <input 
          v-model="searchInput" 
          @keyup.enter="searchStock"
          placeholder="Search by ticker" 
          class="input-primary"
        />
      </div>
      <div class="flex items-center space-x-4">
        <button 
          @click="searchStock" 
          class="btn-primary"
        >
          Search
        </button>
        <button 
          v-if="store.searchTicker" 
          @click="clearSearch" 
          class="btn-secondary"
        >
          Clear
        </button>
      </div>
    </div>

    <div v-if="store.loading" class="text-center py-8">
      <div class="animate-pulse text-gray-500">Loading stocks...</div>
    </div>

    <div v-else-if="store.error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4" role="alert">
      {{ store.error }}
    </div>

    <div v-else-if="store.stocks.length === 0" class="text-center py-8 text-gray-500">
      <p>No stocks found</p>
    </div>

    <div v-else class="space-y-4">
      <div v-if="store.searchTicker" class="bg-blue-50 border-l-4 border-blue-500 p-4 mb-4">
        <p class="text-blue-700">
          Showing results for ticker: 
          <span class="font-bold">{{ store.searchTicker }}</span>
        </p>
      </div>
      
      <StockCard 
        v-for="stock in store.stocks" 
        :key="stock.id" 
        :stock="stock"
      />

      <div 
        v-if="!store.searchTicker && store.pagination.totalPages > 1" 
        class="flex justify-center items-center space-x-4 mt-6"
      >
        <button 
          @click="prevPage" 
          :disabled="store.pagination.currentPage === 1"
          class="btn-secondary"
          :class="{ 'opacity-50 cursor-not-allowed': store.pagination.currentPage === 1 }"
        >
          Previous
        </button>
        <span class="text-gray-600">
          Page {{ store.pagination.currentPage }} of {{ store.pagination.totalPages }}
        </span>
        <button 
          @click="nextPage" 
          :disabled="store.pagination.currentPage === store.pagination.totalPages"
          class="btn-secondary"
          :class="{ 'opacity-50 cursor-not-allowed': store.pagination.currentPage === store.pagination.totalPages }"
        >
          Next
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStocksStore } from '@/stores/stocks';
import StockCard from '@/components/StockCard.vue';

const store = useStocksStore();
const searchInput = ref('');

onMounted(() => {
  store.fetchStocks();
});

const searchStock = () => {
  if (searchInput.value.trim()) {
    store.searchStock(searchInput.value.trim().toUpperCase());
  } else {
    // If search input is empty, revert to normal stocks list
    store.clearSearch();
  }
};

const clearSearch = () => {
  searchInput.value = '';
  store.clearSearch();
};

const prevPage = () => {
  if (store.pagination.currentPage > 1) {
    store.fetchStocks(store.pagination.currentPage - 1);
  }
};

const nextPage = () => {
  if (store.pagination.currentPage < store.pagination.totalPages) {
    store.fetchStocks(store.pagination.currentPage + 1);
  }
};
</script>