package storage

import (
	"reflect"
	"testing"

	apisv1alpha1 "github.com/hwameistor/local-storage/pkg/apis/hwameistor/v1alpha1"
)

func Test_mergeRegistryDiskMap(t *testing.T) {
	type args struct {
		localDiskMap []map[string]*apisv1alpha1.LocalDisk
	}
	var localDiskM []map[string]*apisv1alpha1.LocalDisk
	var localDiskM1 = map[string]*apisv1alpha1.LocalDisk{}
	localDiskM1["/dev/sdb"] = &apisv1alpha1.LocalDisk{
		DevPath:       "/dev/sdb",
		Class:         apisv1alpha1.DiskClassNameHDD,
		CapacityBytes: 10240,
		State:         apisv1alpha1.DiskStateAvailable,
	}
	var localDiskM2 = map[string]*apisv1alpha1.LocalDisk{}
	localDiskM2["/dev/sdc"] = &apisv1alpha1.LocalDisk{
		DevPath:       "/dev/sdc",
		Class:         apisv1alpha1.DiskClassNameHDD,
		CapacityBytes: 102400,
		State:         apisv1alpha1.DiskStateInUse,
	}
	localDiskM = append(localDiskM, localDiskM1)
	localDiskM = append(localDiskM, localDiskM2)
	var wantLocalDiskM = map[string]*apisv1alpha1.LocalDisk{}
	wantLocalDiskM["/dev/sdb"] = &apisv1alpha1.LocalDisk{
		DevPath:       "/dev/sdb",
		Class:         apisv1alpha1.DiskClassNameHDD,
		CapacityBytes: 10240,
		State:         apisv1alpha1.DiskStateAvailable,
	}
	wantLocalDiskM["/dev/sdc"] = &apisv1alpha1.LocalDisk{
		DevPath:       "/dev/sdc",
		Class:         apisv1alpha1.DiskClassNameHDD,
		CapacityBytes: 102400,
		State:         apisv1alpha1.DiskStateInUse,
	}

	tests := []struct {
		name string
		args args
		want map[string]*apisv1alpha1.LocalDisk
	}{
		// TODO: Add test cases.
		{
			args: args{localDiskMap: localDiskM},
			want: wantLocalDiskM,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeRegistryDiskMap(tt.args.localDiskMap...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeRegistryDiskMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
