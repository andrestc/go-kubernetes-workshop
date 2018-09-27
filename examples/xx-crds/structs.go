
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// POD STRUCT OMIT
type Pod struct {
	metav1.TypeMeta   // HL1
	metav1.ObjectMeta // HL1

	Spec   PodSpec   // HL2
	Status PodStatus // HL2
}

// POD STRUCT OMIT

// POD LIST STRUCT OMIT
type PodList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Pod
}

// POD LIST STRUCT OMIT