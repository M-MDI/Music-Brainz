<script type="text/javascript">
function initMap() {
    const locations = [];
    `{{range $in,$arr := .Localisation}}`
    locations.push({ lat: parseFloat('{{index $arr 0}}'), lng: parseFloat('{{index $arr 1}}') });
    `{{end}}`
    const map = new google.maps.Map(document.getElementById("map"), {
        zoom: 2,
        center: locations[0], // Example center point
    });
    locations.forEach(location => {
        new google.maps.Marker({
            position: location,
            map: map,
        });
    });
}
//  key=AIzaSyB8nyzbIc4w06S1g8j7zD7q3JuoDx4FgfE 
document.getElementById("map").initMap = initMap;
</script>
<script src="https://maps.googleapis.com/maps/api/js?key==initMap&v=weekly" defer></script>