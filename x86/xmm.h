#include <stdint.h>
#include <string.h>
#include <xmmintrin.h>

static void xmm_add(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_add_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_sub(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_sub_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_mul(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_mul_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_div(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_div_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_sqrt(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_sqrt_ps(ma);
  _mm_store_ps(out, r);
}

static void xmm_rsqrt(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_rsqrt_ps(ma);
  _mm_store_ps(out, r);
}

static void xmm_rcp(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_rcp_ps(ma);
  _mm_store_ps(out, r);
}

static void xmm_min(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_min_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_max(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_max_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_and(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_and_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_or(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_or_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_xor(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_xor_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_andnot(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_andnot_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_bulk_convert_from_int8(float *out, int8_t *in, int size) {
  int i;
  for(i = 0; i < size; i += 4) {
    __m64 ma = _mm_setr_pi8(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      0, 0, 0, 0
    );
    __m128 r = _mm_cvtpi8_ps(ma);
    _mm_store_ps(out + i, r);
  }
}

static void xmm_bulk_convert_from_uint8(float *out, uint8_t *in, int size) {
  int i;
  for(i = 0; i < size; i += 4) {
    __m64 ma = _mm_setr_pi8(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      0, 0, 0, 0
    );
    __m128 r = _mm_cvtpu8_ps(ma);
    _mm_store_ps(out + i, r);
  }
}

static void xmm_bulk_convert_to_int8(int8_t *out, float *in, int size) {
  int i;
  for(i = 0; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m64 r = _mm_cvtps_pi8(ma);
    memcpy(out + i, &r, sizeof(char) * 4);
  }
}

static void xmm_bulk_convert_to_uint8(uint8_t *out, float *in, int size) {
  int i;
  for(i = 0; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m64 r = _mm_cvtps_pi8(ma);
    // TODO
    r = _mm_slli_pi16(r, 1);
    r = _mm_srai_pi16(r, 1);
    memcpy(out + i, &r, sizeof(unsigned char) * 4);
  }
}

static void xmm_bulk_sum(float *out, float *a, int size) {
  __m128 r = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  int i;
  for(i = 4; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(a[i + 0], a[i + 1], a[i + 2], a[i + 3]);
    r = _mm_add_ps(ma, r);
  }
  _mm_store_ps(out, r);
}

static void xmm_bulk_sum2(float *out, float *a, float *b, int size) {
  int i;
  for(i = 0; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(a[i + 0], a[i + 1], a[i + 2], a[i + 3]);
    __m128 mb = _mm_setr_ps(b[i + 0], b[i + 1], b[i + 2], b[i + 3]);
    __m128 r = _mm_add_ps(ma, mb);
    _mm_store_ps(out + i, r);
  }
}

static void xmm_bulk_add(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  int i;
  for(i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_add_ps(min, mbase);
    _mm_store_ps(out + i, r);
  }
}

static void xmm_bulk_sub(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  int i;
  for(i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_sub_ps(min, mbase);
    _mm_store_ps(out + i, r);
  }
}

static void xmm_bulk_mul(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  int i;
  for(i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_mul_ps(min, mbase);
    _mm_store_ps(out + i, r);
  }
}

static void xmm_bulk_div(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  int i;
  for(i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_div_ps(min, mbase);
    _mm_store_ps(out + i, r);
  }
}

static void xmm_tile4x4_sum(float *out, float *in, int size) {
  int i = 0;
  int p = 0;
  for(; i < size;) {
    __m128 r = _mm_setr_ps(in[i + 0 + 0], in[i + 4 + 0], in[i + 8 + 0], in[i + 12 + 0]); // R
    __m128 g = _mm_setr_ps(in[i + 0 + 1], in[i + 4 + 1], in[i + 8 + 1], in[i + 12 + 1]); // G
    __m128 b = _mm_setr_ps(in[i + 0 + 2], in[i + 4 + 2], in[i + 8 + 2], in[i + 12 + 2]); // B
    __m128 a = _mm_setr_ps(in[i + 0 + 3], in[i + 4 + 3], in[i + 8 + 3], in[i + 12 + 3]); // A
    
    __m128 rg = _mm_add_ps(r, g);
    __m128 ba = _mm_add_ps(b, a);
    __m128 gray = _mm_add_ps(rg, ba);
    _mm_stream_ps(out + p, gray);

    i += 16;
    p += 4;
  }
}
