package place

import (
	"context"
	"strconv"

	"github.com/NetKBs/backend-reviewapp/geoapify"
)

func GetPlaceDetailsByMapsIdService(ctx context.Context, mapsID string) (placeDetailsDTO PlaceDetailsResponseDTO, err error) {
	placeDetails, err := geoapify.GetPlaceDetailsById(mapsID)
	if err != nil {
		return placeDetailsDTO, err
	}
	place, err := findPlaceByMapsIdRepo(mapsID)
	if err != nil {
		return placeDetailsDTO, err
	}

	placeDetailsDTO = PlaceDetailsResponseDTO{
		ID:        place.ID,
		Details:   placeDetails,
		CreatedAt: place.CreatedAt.String(),
		UpdatedAt: place.UpdatedAt.String(),
	}
	return placeDetailsDTO, err
}

func GetPlaceDetailsByPlaceIdService(ctx context.Context, placeId int) (placeDetailsDTO PlaceDetailsResponseDTO, err error) {

	place, err := findPlaceByPlaceIdRepo(placeId)
	if err != nil {
		return placeDetailsDTO, err
	}
	placeDetails, err := geoapify.GetPlaceDetailsById(place.MapsId)
	if err != nil {
		return placeDetailsDTO, err
	}

	placeDetailsDTO = PlaceDetailsResponseDTO{
		ID:        place.ID,
		Details:   placeDetails,
		CreatedAt: place.CreatedAt.String(),
		UpdatedAt: place.UpdatedAt.String(),
	}
	return placeDetailsDTO, err
}

func GetPlaceDetailsByCoordsService(ctx context.Context, lon, lat string) (placeDetailsDTO PlaceDetailsResponseDTO, err error) {
	placeDetails, err := geoapify.GetPlaceDetailsByCoord(lon, lat)
	if err != nil {
		return placeDetailsDTO, err
	}
	place, err := findPlaceByMapsIdRepo(placeDetails.MapsID)
	if err != nil {
		return placeDetailsDTO, err
	}

	placeDetailsDTO = PlaceDetailsResponseDTO{
		ID:        place.ID,
		Details:   placeDetails,
		CreatedAt: place.CreatedAt.String(),
		UpdatedAt: place.UpdatedAt.String(),
	}
	return placeDetailsDTO, err
}

// categories: slice of strings
func GetPlacesByCoordsService(ctx context.Context, categories []string, lon, lat string) (places PlacesResponseDTO, err error) {
	var catString string
	for i, v := range categories {
		if i == 0 {
			catString = v
		} else {
			catString = catString + "," + v
		}
	}
	placesRaw, err := geoapify.GetPlacesAroundCoords(catString, lon, lat)

	if err != nil {
		return places, err
	}

	centerLon, _ := strconv.ParseFloat(lon, 64)
	centerLat, _ := strconv.ParseFloat(lat, 64)
	places = PlacesResponseDTO{
		CenterLon: centerLon,
		CenterLan: centerLat,
		Data:      placesRaw,
	}
	return places, err
}

func GetAutocompleteResultService(ctx context.Context, text string) (autocompleteResult AutocompleteResponseDTO, err error) {
	geocodings, err := geoapify.GetAutocompleteResponse(text)
	autocompleteResult.Query = text
	autocompleteResult.Result = geocodings

	return autocompleteResult, err
}
