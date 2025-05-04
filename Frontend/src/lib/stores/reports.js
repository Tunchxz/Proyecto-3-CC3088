import { writable } from 'svelte/store';

// Estado inicial
const initialState = {
  reports: [],
  filters: {
    dateFrom: '',
    dateTo: '',
    reportType: '',
    status: ''
  },
  loading: false,
  error: null
};

function createReportsStore() {
  const { subscribe, set, update } = writable(initialState);

  return {
    subscribe,
    setReports: (reports) => update(state => ({ ...state, reports })),
    setFilters: (filters) => update(state => ({ ...state, filters })),
    setLoading: (loading) => update(state => ({ ...state, loading })),
    setError: (error) => update(state => ({ ...state, error })),
    reset: () => set(initialState)
  };
}

export const reports = createReportsStore(); 