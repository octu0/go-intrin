#include <string.h>
#include <xmmintrin.h>

static void xmm_add(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_add_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_sub(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_sub_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_mul(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_mul_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_div(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_div_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_sqrt(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_sqrt_ps(ma);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_rsqrt(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_rsqrt_ps(ma);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_rcp(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_rcp_ps(ma);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_min(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_min_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_max(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_max_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_and(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_and_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_or(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_or_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_xor(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_xor_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}

static void xmm_andnot(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_andnot_ps(ma, mb);
  memcpy(out, &r, sizeof(float) * 4);
}
