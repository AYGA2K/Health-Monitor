<template>
  <div class="q-ma-xl">
    <div id="chart" class="q-px-lg">
      <apexchart type="line" height="400" :key="componentKey" :options="chartOptions" :series="chartSeries"></apexchart>
    </div>
  </div>
</template>

<script lang="ts">

import { state } from "./../socketGo";



export default {
  computed: {
    chartSeries() {
      interface Message {
        heartbeat: string,
        time: string
      }
      return [{
        name: 'series-1',
        data: state.data.map((message: Message) => message.heartbeat)
        //data: state.data.map((message) => message.heartbeat).slice(-5)
      }]
    },
    chartOptions() {
      interface Message {
        heartbeat: string,
        time: string
      }
      return {

        chart: {
          animations: {
            enabled: true,
            easing: 'easeinout',
            speed: 800,
            animateGradually: {
              enabled: true,
              delay: 150
            },
            dynamicAnimation: {
              enabled: true,
              speed: 350
            }
          },

          type: 'line',
        },
        dataLabels: {
          enabled: true,
        },
        stroke: {
          curve: 'smooth',
        },
        title: {
          text: 'heartbeat',
          align: 'left',
          style: {
            fontSize: '14px',
            fontWeight: 'bold',

          },
        },
        tooltip: { enabled: true },
        grid: {
          show: false,
        },
        xaxis: {
          categories: state.data.map((message: Message) => message.time),
          // categories: state.data.map((message) => message.time).slice(-5),

          title: {
            text: 'Time',
            style: {
              fontSize: '14px',
              fontWeight: 'bold',

            },
          },
          labels: {
            show: false,
            rotate: -90,
            rotateAlways: true,
            hideOverlappingLabels: true,
            trim: true
          }
        }
      }

    }
  },
  data() {
    return {
      componentKey: 0,
    };
  },
  methods: {
    forceRerender() {
      this.componentKey += 1;
    }

  },

}
</script>
