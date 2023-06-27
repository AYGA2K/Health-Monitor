<template>
  <div class="q-pa-sm column q-gutter-md ">
    <div class="row justify-around items-center ">
      <q-card class=" border-radius-10px col-12  column " style="max-width: 500px;">
        <q-card-section class="bg-blue-grey-14 col-4  ">
          <div class="text-center text-white text-h5">Today</div>
          <div class="row justify-around">
            <div class=" column items-center  q-pa-lg circle-1">
              <div class="col">
                <q-icon size="sm" name="fa-solid  fa-person-walking" color="white" />

              </div>
              <div class="col text-white  text-h4">
                {{ steps }}
              </div>
              <div class="col text-blue-1 " style="padding: 1rem 0 0 0;">
                Steps
              </div>
            </div>
            <div class=" column items-center justify-around q-pa-lg circle-1">
              <div class="col">
                <q-icon size="sm" name="alarm" color="white" />

              </div>
              <div class="col text-white  text-h4">
                0
              </div>
              <div class="col text-blue-1 " style="padding: 1rem 0 0 0;">
                Min
              </div>
            </div>

          </div>

        </q-card-section>
        <q-card-section class=" bg-blue-grey-9 col-4  ">
          <div class="row justify-around">
            <div class=" text-white">
              <div class="column items-center">
                <div class="col"> Distance </div>
                <div class="col"> 1.2 Km </div>
              </div>
            </div>

            <div class=" text-white">
              <div class="column items-center">
                <div class="col"> Calories </div>
                <div class="col"> 23 Kcal </div>
              </div>
            </div>

          </div>

        </q-card-section>


      </q-card>

    </div>
    <div>
      <q-card class=" section-2 row bg-transparent text-white  q-gutter-sm no-box-shadow justify-between content-center ">
        <q-card-section class="col text-center  bg-teal-10">
          <div class="text-h6">Sleep</div>
          <q-icon size="md" color="orange" name="hotel"> </q-icon>
          <div class="q-mt-lg">{{ sleep }} Hours</div>
        </q-card-section>
        <q-card-section class="col text-center  bg-indigo-10 rounded-borders">
          <div class="text-h6">Weight</div>
          <q-icon size="md" color="red-5" name="monitor_weight"> </q-icon>
          <div class="q-mt-lg">70 Kg</div>
        </q-card-section>
        <q-card-section class="col text-center bg-red-10 rounded-borders">
          <div class="text-h6">Heart rate</div>
          <q-icon size="md" color="cyan-4" name="monitor_heart"> </q-icon>
          <div class="q-mt-lg">85 bpm</div>
        </q-card-section>
      </q-card>


    </div>
    <div id="chart" class="q-px-lg">
      <apexchart type="line" height="300" :options="chartOptions" :series="chartSeries"></apexchart>
    </div>
  </div>
</template>
<script lang="ts" >

import { heartbeats_state } from "src/services/heartbeat";
import { sleep_state } from "src/services/sleep";
import { step_state } from "src/services/step";
import { ip } from 'src/ip_adress';
interface Message {
  value: number,
  CreatedAt: string
}
export default {
  computed: {
    steps() {
      return step_state.data.map((message: Message) => message.value).slice(-1)[0]
    },
    sleep() {
      return sleep_state.data.map((message: Message) => message.value).slice(-1)[0]

    },
    chartSeries() {
      interface Message {
        value: number,
        CreatedAt: string
      }
      return [{
        name: 'heartbeats',
        data: heartbeats_state.data.map((message: Message) => message.value)
        //data: state.data.map((message) => message.heartbeat).slice(-5)
      }]
    },
    chartOptions() {
      interface Message {
        value: number,
        CreatedAt: string
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
          categories: heartbeats_state.data.map((message: Message) => message.CreatedAt),
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
    };
  },
  methods: {

  },

}


</script>
<style lang="scss">
.text-subtitle2 {
  color: dark;
}

.border-radius-10px {
  border-radius: 10px !important;
  border: 2px;
}

.bg-primary {
  background: $primary ;
}


.circle-1 {
  min-width: 130px;
  border-radius: 50% !important;
  box-shadow:
    0 0 0 10px $blue-3,
}

.section-2 .q-card__section {
  border-radius: 5% !important;
}
</style>
