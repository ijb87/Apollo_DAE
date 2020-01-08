<template>
  <div>
    <input type="text" class="coupon-code" v-model="code" />
    <p v-text="feedback"></p>
  </div>
</template>

<script>
export default {
  data () {
    return {
      code: '',
      coupons: [
        {
          code: '10OFF',
          message: '10% Off!',
          discount: 10
        },
        {
          code: 'FREE',
          message: 'Entirely Free!',
          discount: 100
        }
      ],
      valid: false
    }
  },
  computed: {
    selectedCoupon () {
      this.coupons.find(coupon => coupon.code == this.code)
    },
    message () {
      return this.selectedCoupon.message
    },

    feedback () {
      if (this.valid) {
        return `Coupon Redeemed: ${this.message}`
      }

      return 'Invalid Coupon Code'
    }
  },
  methods: {
    validate () {
      this.valid = !!this.selectedCoupon

      if (this.valid) {
        this.$emit('applied', this.selectedCoupon().discount)
      }
    }
  }
}
</script>
