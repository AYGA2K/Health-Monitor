<template>
  <div id="chart">
    <apexchart type="line" height="190" :options="chartOptions" :series="series"></apexchart>
  </div>
</template>

<script>
export default {
  name: 'myChart2',
  data() {
    return {
      series: [
        {
          name: 'Desktops',
          data: [41, 35, 51, 49, 62, 69, 91],
        },
      ],
      chartOptions: {
        chart: {
          height: 350,
          type: 'line',
          zoom: {
            enabled: false,
          },
        },
        dataLabels: {
          enabled: false,
        },
        stroke: {
          curve: 'straight',
        },
        title: {
          text: 'Product Trends by Month',
          align: 'left',
          style: {
            fontSize: '14px',
            fontWeight: 'bold',

            color: '#fff  ',
          },
        },
        tooltip: { enabled: false },
        grid: {
          show: false,
        },
        xaxis: {
          categories: [
            'Jan',
            'Feb',
            'Mar',
            'Apr',
            'May',
            'Jun',
            'Jul',
          ],
        },
      },
    };
  },
};
</script>

<style lang="scss"></style>
