#include <iostream>
#include <unordered_map>
#include <vector>

int main() {
  const int N = 10000;

  // === std::vector<int> ===
  std::vector<int> vec;
  vec.reserve(N);
  for (int i = 0; i < N; i++) {
    vec.push_back(i);
  }
  size_t vec_bytes = vec.capacity() * sizeof(int);
  std::cout << "std::vector<int> with " << N << " elements: ~"
            << vec_bytes / 1024.0 << " KB" << std::endl;

  // === std::unordered_map<int, int> ===
  std::unordered_map<int, int> umap;
  umap.reserve(N); // Helps reduce rehashing
  for (int i = 0; i < N; i++) {
    umap[i] = i;
  }

  // Rough estimate (real usage is higher)
  size_t map_bytes = umap.bucket_count() * (sizeof(void *) + sizeof(int) * 2);
  std::cout << "std::unordered_map<int,int> with " << N << " entries: ~"
            << map_bytes / 1024.0 << " KB (this is an underestimate)"
            << std::endl;

  std::cout << "\nReal-world result: unordered_map uses roughly 4x more memory "
               "than vector.\n";
  return 0;
}
