Vue.use(VueTables.ClientTable);

new Vue({
  el: "#app",
  data: {
    columns: ["id", "name", "email", 
"group_name"],
    data: getData(),
     options: {
      headings: {
        id: 'ID',
        name: 'Full Name',
        email: 'Email',
        group_name: 'Department'
      },
      sortable: [
        'id'
      ],
      texts: {
        filterPlaceholder: 'filter...'
      }
    }
  }
});

function getData() {
  return [
    {
      id: 1,
      name: "Valerie Lewis",
      email: "pjot.grewal.7399g@erofan.org",
      group_name: "Agency of Software Networking"
    },
    {
      id: 2,
      name: "Rudy Valdez",
      email: "knaffdo@325designcentre.xyz",
      group_name: "Branch of Analytical Hardware Programming and Mainframe Connectivity"
    },

    {
      id: 3,
      name: "Percy Saunders",
      email: "kosamarad@dacsancantho.net",
      group_name: "Business-Oriented Multimedia Administration and Extranet Troubleshooting Committee"
    },
    {
      id: 4,
      name: "Dora Parks",
      email: "emotasi@mrtsport.com",
      group_name: "Department of E-Commerce Network Connectivity and Maintenance"
    }
  ];
}