package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PetType is the type of the pet. For example, a dog.
// +kubebuilder:validation:Enum=dog;cat;bird;reptile
type PetType string

const (
	DogPetType     PetType = "dog"
	CatPetType     PetType = "cat"
	BirdPetType    PetType = "bird"
	ReptilePetType PetType = "reptile"
)

// PetSpec defines the desired state of Pet
type PetSpec struct {
	// Name is the name of the pet
	Name string `json:"name"`
	// Type is the type of pet
	Type PetType `json:"type"`
	// Birthday is the date the pet was born
	Birthday metav1.Time `json:"birthday"`
}

// PetStatus defines the observed state of Pet
type PetStatus struct {
	// ID is the unique identifier created by the service for the pet
	ID string `json:"id,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pet is the Schema for the pets API
type Pet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PetSpec   `json:"spec,omitempty"`
	Status PetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PetList contains a list of Pet
type PetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pet{}, &PetList{})
}